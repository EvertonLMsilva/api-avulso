package controller

import (
	"fmt"
	"net/http"
)

// GetAllUser list all users
func GetAllUser(reqW http.ResponseWriter, res http.Request) {
	resolveJson := `{ 
		id: 00,
		name: "Everton"
	 }`
	fmt.Fprintf(reqW, resolveJson)
}
