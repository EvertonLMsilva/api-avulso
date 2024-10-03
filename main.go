package main

import (
	"fmt"
	"log"
	"net/http"
)

func PortServer() (res string) {
	const PORT int32 = 3000
	return fmt.Sprintf(":%v", PORT)
}

func main() {
	http.HandleFunc("/user/all", UserGetAllDto)

	log.Println("Server listennig on $s", PortServer())
	log.Fatal(http.ListenAndServe(PortServer(), nil))
}
