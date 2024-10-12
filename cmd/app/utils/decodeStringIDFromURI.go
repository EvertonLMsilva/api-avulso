package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func DecodeStringIDFromURI(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		fmt.Println("Id Disable 02", id)
		return "", errors.New("empty_id_error")
	}
	return id, nil
}
