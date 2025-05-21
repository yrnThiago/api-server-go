package usecase

import (
	"context"
	"encoding/json"
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/entity"
	infra "github.com/yrnThiago/api-server-go/internal/infra/redis"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/user"
	"github.com/yrnThiago/api-server-go/internal/utils"
	"go.uber.org/zap"
)

type AuthInputDto struct {
	Email    string
	Password string
}

type AuthUseCase struct {
	UserRepository usecase.IUserRepository
}

func NewAuthUseCase(userRepository usecase.IUserRepository) *AuthUseCase {
	return &AuthUseCase{
		UserRepository: userRepository,
	}
}

func (a *AuthUseCase) Login(input AuthInputDto) (string, *entity.User, error) {

	output, err := a.GetByLogin(input)
	if err != nil {
		return "", nil, err
	}

	err = utils.CheckPasswordHash(input.Password, output.Password)
	if err != nil {
		config.Logger.Warn(entity.ErrWrongCredentialsMsg)
		return "", nil, err
	}

	token, err := utils.GenerateJWT(output.ID)
	if err != nil {
		config.Logger.Fatal(
			"jwt token not generated",
			zap.Error(err),
		)

		return "", nil, err
	}

	userJson, err := json.Marshal(output)
	if err != nil {
		config.Logger.Fatal(
			"marshal user json",
			zap.Error(err),
		)

		return "", nil, err
	}

	infra.Redis.Set(context.Background(), "user-"+output.ID, string(userJson), config.Env.UserSessionExpiresAt)
	return token, output, nil
}

func (a *AuthUseCase) Logout() error {
	return nil
}

func (a *AuthUseCase) GetByLogin(authInputDto AuthInputDto) (*entity.User, error) {
	user, err := a.UserRepository.GetByLogin(authInputDto.Email)
	if err != nil {
		return nil, err
	}

	return user, err
}
