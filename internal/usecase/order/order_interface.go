package usecase

import "github.com/yrnThiago/api-server-go/internal/entity"

type IOrderRepository interface {
	Add(order *entity.Order) error
	GetMany() ([]*entity.Order, error)
	GetById(id string) (*entity.Order, error)
	UpdateById(order, newOrderBody *entity.Order) (*entity.Order, error)
	DeleteById(id string) error
}
