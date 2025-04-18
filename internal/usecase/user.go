package usecase

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

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
	email string
	gorm.Model
}

type UserUseCase struct {
	UserRepository entity.UserRepository
}

func NewUserUseCase(userRepository entity.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Add(
	input UserInputDto,
) (*entity.User, error) {
	validationError := utils.ValidateStruct(input)
	if validationError != nil {
		return nil, utils.NewErrorInfo("ValidationError", validationError.Error())
	}

	input.Password, _ = utils.GenerateHashPassword(input.Password)

	user := entity.NewUser(input.Email, input.Password)
	err := u.UserRepository.Add(user)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("new user added")
	return user, err
}

func (u *UserUseCase) GetMany() ([]*entity.User, error) {
	users, err := u.UserRepository.GetMany()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUseCase) GetById(id string) (*entity.User, error) {
	user, err := u.UserRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) GetByEmail(email string) (*entity.User, error) {
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) UpdateById(
	id string,
	input UserInputDto,
) (*entity.User, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, utils.NewErrorInfo("ValidationError", err.Error())
	}

	user, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	newUserBody := entity.NewUser(input.Email, input.Password)
	updatedUser, err := u.UserRepository.UpdateById(user, newUserBody)
	if err != nil {
		return nil, err
	}

	config.Logger.Info(
		"user updated",
		zap.String("id", id),
	)
	return updatedUser, nil
}

func (u *UserUseCase) DeleteById(
	id string,
) error {
	err := u.UserRepository.DeleteById(id)
	if err != nil {
		return err
	}

	config.Logger.Info("user deleted",
		zap.String("id", id),
	)
	return err
}
