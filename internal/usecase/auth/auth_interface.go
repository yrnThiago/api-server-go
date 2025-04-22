package usecase

import "github.com/yrnThiago/api-server-go/internal/entity"

type IAuthUseCase interface {
	Login(input AuthInputDto) (string, *entity.User, error)
	Logout() error
	GetByLogin(authInputDto AuthInputDto) (*entity.User, error)
}
