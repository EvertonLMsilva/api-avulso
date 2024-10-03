package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/all", GetAllUser)

	port := 3000
	log.Println("Server listennig on $s", port)
}
