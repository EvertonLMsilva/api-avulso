package useCase

import "github.com/EvertonLMsilva/api-avulso/internal/entity"

type ListUsersOutputDto struct {
	ID       string
	Name     string
	Birthday string
	Active   bool
}

type ListUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewListUserUseCase(userRepository entity.UserRepository) *ListUserUseCase {
	return &ListUserUseCase{UserRepository: userRepository}
}

func (u *ListUserUseCase) Execute() ([]*ListUsersOutputDto, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var userOutput []*ListUsersOutputDto
	for _, user := range users {
		userOutput = append(userOutput, &ListUsersOutputDto{
			ID:       user.ID,
			Name:     user.Name,
			Birthday: user.Birthday,
			Active:   user.Active,
		})
	}

	return userOutput, nil
}
