package usecase

import "github.com/yrnThiago/api-server-go/internal/entity"

type IProductRepository interface {
	Add(product *entity.Product) error
	GetMany() ([]*entity.Product, error)
	GetById(id string) (*entity.Product, error)
	UpdateById(product *entity.Product) (*entity.Product, error)
	DeleteById(id string) error
}
