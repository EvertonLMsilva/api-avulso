package web

import (
	"encoding/json"
	"net/http"

	"github.com/EvertonLMsilva/api-avulso/cmd/app/utils"
	"github.com/EvertonLMsilva/api-avulso/internal/useCase"
)

type UserHandlers struct {
	CreateUserUseCase  *useCase.CreateUserUseCase
	ListUserUseCase    *useCase.ListUserUseCase
	DisableUserUseCase *useCase.DisableUserUseCase
}

func NewUserHandlers(
	createUserUseCase *useCase.CreateUserUseCase,
	listUserUseCase *useCase.ListUserUseCase,
	disableUseCase *useCase.DisableUserUseCase,
) *UserHandlers {
	return &UserHandlers{
		CreateUserUseCase:  createUserUseCase,
		ListUserUseCase:    listUserUseCase,
		DisableUserUseCase: disableUseCase,
	}
}

func (u *UserHandlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input useCase.CreateUserInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(err.Error()))
		return
	}

	output, err := u.CreateUserUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
	return
}

func (u *UserHandlers) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	output, err := u.ListUserUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (u *UserHandlers) DisableUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.DecodeStringIDFromURI(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	errFormatId := u.DisableUserUseCase.Execute(id)
	if errFormatId != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
