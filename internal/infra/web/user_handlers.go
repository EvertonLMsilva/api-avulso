package web

import (
	"encoding/json"
	"net/http"

	"github.com/EvertonLMsilva/api-avulso/internal/usecase"
)

type UserHandlers struct {
	CreateUserUseCase *usecase.CreateUserUseCase
	ListUserUseCase   *usecase.ListUserUseCase
}

func NewUserHandlers(createUserUseCase *usecase.CreateUserUseCase, listUserUseCase *usecase.ListUserUseCase) *UserHandlers {
	return &UserHandlers{
		CreateUserUseCase: createUserUseCase,
		ListUserUseCase:   listUserUseCase,
	}
}

func (u *UserHandlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	output, err := u.CreateUserUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (u *UserHandlers) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	output, err := u.ListUserUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
