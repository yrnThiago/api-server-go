package usecase

import (
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type UserInputDto struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

type UserOutputDto struct {
	ID    string
	Email string
}

type UserUseCase struct {
	UserRepository IUserRepository
}

func NewUserUseCase(userRepository IUserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Add(
	input UserInputDto,
) (*UserOutputDto, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, entity.GetValidationError(err.Error())
	}

	input.Password, _ = utils.GenerateHashPassword(input.Password)

	user := entity.NewUser(input.Email, input.Password)
	err = u.UserRepository.Add(user)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("new user added")
	return &UserOutputDto{
		ID:    user.ID,
		Email: user.Email,
	}, err
}

func (u *UserUseCase) GetMany() ([]*UserOutputDto, error) {
	var usersOutputDTO []*UserOutputDto
	users, err := u.UserRepository.GetMany()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userOutputDTO := &UserOutputDto{
			ID:    user.ID,
			Email: user.Email,
		}

		usersOutputDTO = append(usersOutputDTO, userOutputDTO)
	}

	return usersOutputDTO, nil
}

func (u *UserUseCase) GetById(id string) (*UserOutputDto, error) {
	user, err := u.UserRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDto{
		ID:    user.ID,
		Email: user.Email,
	}, err
}

func (u *UserUseCase) GetByLogin(email string) (*entity.User, error) {
	user, err := u.UserRepository.GetByLogin(email)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *UserUseCase) UpdateById(
	id string,
	input UserInputDto,
) (*UserOutputDto, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, entity.GetValidationError(err.Error())
	}

	user, err := u.UserRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	user.Email = input.Email
	user.Password = input.Password

	updatedUser, err := u.UserRepository.UpdateById(user)
	if err != nil {
		return nil, err
	}

	config.Logger.Info(
		"user updated",
		zap.String("id", id),
	)

	return &UserOutputDto{
		ID:    updatedUser.ID,
		Email: updatedUser.Email,
	}, err
}

func (u *UserUseCase) DeleteById(
	id string,
) (*UserOutputDto, error) {
	user, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.DeleteById(id)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("user deleted",
		zap.String("id", id),
	)
	return user, err
}
