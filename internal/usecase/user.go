package usecase

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/models"
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
	UserRepository models.UserRepository
}

func NewUserUseCase(userRepository models.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Add(
	input UserInputDto,
) (*models.User, error) {
	validationError := utils.ValidateStruct(input)
	if validationError != nil {
		return nil, utils.NewErrorInfo("ValidationError", validationError.Error())
	}

	input.Password, _ = utils.GenerateHashPassword(input.Password)

	user := models.NewUser(input.Email, input.Password)
	err := u.UserRepository.Add(user)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("new user added")
	return user, err
}

func (u *UserUseCase) GetMany() ([]*models.User, error) {
	users, err := u.UserRepository.GetMany()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUseCase) GetById(id string) (*models.User, error) {
	user, err := u.UserRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) GetByEmail(email string) (*models.User, error) {
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) UpdateById(
	id string,
	input UserInputDto,
) (*models.User, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, utils.NewErrorInfo("ValidationError", err.Error())
	}

	user, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	newUserBody := models.NewUser(input.Email, input.Password)
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
