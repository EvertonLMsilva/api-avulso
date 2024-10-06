package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string
	Name     string
	Birthday string
}

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]*User, error)
}

func NewUser(name string, birthday string) *User {
	return &User{
		ID:       uuid.New().String(),
		Name:     name,
		Birthday: birthday,
	}
}
