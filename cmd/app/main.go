package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EvertonLMsilva/api-avulso/internal/infra/akafka"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/repository"
	"github.com/EvertonLMsilva/api-avulso/internal/infra/web"
	"github.com/EvertonLMsilva/api-avulso/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
)

func PortServer() (res string) {
	const PORT int32 = 3000
	return fmt.Sprintf(":%v", PORT)
}

func main() {

	db, err := sql.Open(PostgresDriver, DataSourceName)
	if err != nil {
		panic(err)
	}

	db.Close()

	repository := repository.NewUserRepositoryPg(db)
	createUserUseCase := usecase.NewCreateUserUseCase(repository)
	listUserUseCase := usecase.NewListUserUsecase(repository)

	userHandler := web.NewUserHandlers(createUserUseCase, listUserUseCase)

	r := chi.NewRouter()
	r.Post("/user", userHandler.CreateUserHandler)
	r.Get("/user", userHandler.ListUserHandler)

	log.Println("Server listennig on $s", PortServer())
	go http.ListenAndServe(PortServer(), r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"users"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateUserInputDto{}
		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			//log de erro
			fmt.Printf("Message kafka error", err)
		}
		_, err = createUserUseCase.Execute(dto)
	}

	// http.HandleFunc("/user/all", controllers.GetAllUser)

	// log.Println("Server listennig on $s", PortServer())
	// log.Fatal(http.ListenAndServe(PortServer(), nil))
}
