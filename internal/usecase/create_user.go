package usecase

import "github.com/EvertonLMsilva/api-avulso/internal/entity"

type CreateUserInputDto struct {
	Name     string "json:name"
	Birthday string "json:birthday"
}

type CreateUserOutputDto struct {
	ID       string
	Name     string
	Birthday string
}

type CreateUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewCreateUserUseCase(userRepository entity.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: userRepository}
}

func (u *CreateUserUseCase) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {
	user := entity.NewUser(input.Name, input.Birthday)
	err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDto{
		ID:       user.ID,
		Name:     user.Name,
		Birthday: user.Birthday,
	}, nil
}
