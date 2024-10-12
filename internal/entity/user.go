package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string
	Name     string
	Birthday string
	Active   bool
}

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]*User, error)
	Update(id string, user *User) (*User, error)
	Disable(id string) error
}

func NewUser(name string, birthday string, active bool) *User {
	return &User{
		ID:       uuid.New().String(),
		Name:     name,
		Birthday: birthday,
		Active:   active,
	}
}

func UpdateUser(name string, birthday string, active bool) *User {
	return &User{
		Name:     name,
		Birthday: birthday,
		Active:   active,
	}
}
