package usecase

import "github.com/yrnThiago/api-server-go/internal/entity"

type IUserRepository interface {
	Add(user *entity.User) (*entity.User, error)
	GetMany() ([]*entity.User, error)
	GetById(id string) (*entity.User, error)
	GetByLogin(email string) (*entity.User, error)
	UpdateById(user *entity.User) (*entity.User, error)
	DeleteById(id string) error
}
