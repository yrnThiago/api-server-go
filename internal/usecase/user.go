package usecase

import (
	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/models"
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
	user := models.NewUser(input.Email, input.Password)
	err := u.UserRepository.Add(user)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("adding new user")
	return user, nil
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
	userId string,
	input *UserInputDto,
) (*models.User, error) {
	newUser := models.NewUser(input.Email, input.Password)
	user, err := u.UserRepository.UpdateById(userId, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) DeleteById(
	userId string,
) error {
	err := u.UserRepository.DeleteById(userId)
	if err != nil {
		return err
	}

	return err
}
