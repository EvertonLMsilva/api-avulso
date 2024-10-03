package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAllUser list all users
func GetAllUser(w http.ResponseWriter, r *http.Request) (resUser UserGetAllDto) {
	resUser = UserGetAllDto{
		name:     "everton",
		birthday: "23/12/1990",
		email:    "everton.l.m.silva@gmail.com",
	}

	b, err := json.Marshal(resUser)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	return resUser
}
