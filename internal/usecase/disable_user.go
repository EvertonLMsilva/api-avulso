package useCase

import (
	"github.com/EvertonLMsilva/api-avulso/internal/entity"
)

type DisableUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewDisableUserUseCase(userRepository entity.UserRepository) *DisableUserUseCase {
	return &DisableUserUseCase{UserRepository: userRepository}
}

func (u *DisableUserUseCase) Execute(id string) error {
	err := u.UserRepository.Disable(id)
	if err != nil {
		return err
	}

	return nil
}
