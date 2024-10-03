package user

import (
	"fmt"
	"net/http"
)

// GetAllUser list all users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	resolveJson := `{ 
		id: 00,
		name: "Everton"
	 }`
	fmt.Fprintf(w, resolveJson)
}
