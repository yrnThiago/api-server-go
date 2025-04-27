package usecase

import (
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type UserUseCase struct {
	UserRepository IUserRepository
}

func NewUserUseCase(userRepository IUserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}
func (u *UserUseCase) Add(
	input dto.UserInputDto,
) (*dto.UserOutputDto, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	input.Password, _ = utils.GenerateHashPassword(input.Password)

	user := entity.NewUser(input.Email, input.Password)
	err = u.UserRepository.Add(user)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("new user added")
	return dto.NewUserOutputDto(user), nil
}

func (u *UserUseCase) GetMany() ([]*dto.UserOutputDto, error) {
	var usersOutputDTO []*dto.UserOutputDto
	users, err := u.UserRepository.GetMany()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		usersOutputDTO = append(usersOutputDTO, dto.NewUserOutputDto(user))
	}

	return usersOutputDTO, nil
}

func (u *UserUseCase) GetById(id string) (*dto.UserOutputDto, error) {
	user, err := u.UserRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewUserOutputDto(user), nil
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
	input dto.UserInputDto,
) (*dto.UserOutputDto, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
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

	return dto.NewUserOutputDto(updatedUser), nil
}

func (u *UserUseCase) DeleteById(
	id string,
) (*dto.UserOutputDto, error) {
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
