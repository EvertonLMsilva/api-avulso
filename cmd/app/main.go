package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EvertonLMsilva/api-avulso/cmd/app/dbConfig"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/repository"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/web"
	"github.com/EvertonLMsilva/api-avulso/internal/useCase"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func PortServer() (res string) {
	const PORT int32 = 3000
	return fmt.Sprintf(":%v", PORT)
}

func main() {
	db := dbConfig.ConnectDb()
	defer db.Close()

	repository := repository.NewUserRepositoryPg(db)

	createUserUseCase := useCase.NewCreateUserUseCase(repository)
	listUserUseCase := useCase.NewListUserUseCase(repository)
	disableUserUseCase := useCase.NewDisableUserUseCase(repository)

	userHandler := web.NewUserHandlers(createUserUseCase, listUserUseCase, disableUserUseCase)

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
		router.Delete("/user/{id}", userHandler.DisableUserHandler)
	})

	log.Printf("Server listening on %s", PortServer())
	http.ListenAndServe(PortServer(), r)
}
