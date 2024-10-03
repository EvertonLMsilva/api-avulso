package main

import (
	"log"
	"net/http"

	user "github.com/EvertonLMsilva/api-avulso/pkg/controller/user"
)

func main() {
	http.HandleFunc("/user/all", user.GetAllUser)

	port := 3000
	log.Println("Server listennig on $s", port)
}
