package useCase

import (
	"github.com/EvertonLMsilva/api-avulso/internal/entity"
)

type UpdateUserInputDto struct {
	Name     string "json:name"
	Birthday string "json:birthday"
	Active   bool   "json:active"
}

type UpdateUserOutputDto struct {
	ID       string
	Name     string
	Birthday string
	Active   bool
}

type UpdateUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewUpdateUserUseCase(userRepository entity.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepository: userRepository}
}

func (u *UpdateUserUseCase) Execute(input UpdateUserInputDto, id string) (*UpdateUserOutputDto, error) {
	user := entity.UpdateUser(input.Name, input.Birthday, input.Active)
	resolveUpdate, err := u.UserRepository.Update(id, user)

	if err != nil {
		return nil, err
	}

	return &UpdateUserOutputDto{
		ID:       resolveUpdate.ID,
		Name:     resolveUpdate.Name,
		Birthday: resolveUpdate.Birthday,
		Active:   resolveUpdate.Active,
	}, nil
}
