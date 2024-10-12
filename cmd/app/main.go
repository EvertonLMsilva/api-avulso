package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/EvertonLMsilva/api-avulso/cmd/app/utils"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/dbConfig"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/environments"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/repository"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/web"
	"github.com/EvertonLMsilva/api-avulso/internal/useCase"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	err := environments.StartConfig()
	utils.FatalError(err)

	db := dbConfig.ConnectDb()
	defer db.Close()

	repository := repository.NewUserRepositoryPg(db)

	createUserUseCase := useCase.NewCreateUserUseCase(repository)
	listUserUseCase := useCase.NewListUserUseCase(repository)
	disableUserUseCase := useCase.NewDisableUserUseCase(repository)
	updateUserUseCase := useCase.NewUpdateUserUseCase(repository)

	userHandler := web.NewUserHandlers(createUserUseCase, listUserUseCase, disableUserUseCase, updateUserUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	if r == nil {
		panic(r)
	}

	r.Route("/api", func(router chi.Router) {
		router.Post("/user", userHandler.CreateUserHandler)
		router.Get("/user", userHandler.ListUserHandler)
		router.Put("/user/{id}", userHandler.UpdateUserHandler)
		router.Delete("/user/{id}", userHandler.DisableUserHandler)
	})

	log.Printf("Server listening on %d", environments.Env.ApiPort)

	s := &http.Server{
		Addr:           PortServer(),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

func PortServer() (res string) {
	return fmt.Sprintf(":%d", environments.Env.ApiPort)
}
