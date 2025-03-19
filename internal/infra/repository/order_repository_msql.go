package repository

import (
	"gorm.io/gorm"

	"github.com/yrnThiago/gdlp-go/internal/domain"
)

type OrderRepositoryMysql struct {
	DB *gorm.DB
}

func NewOrderRepositoryMysql(db *gorm.DB) *OrderRepositoryMysql {
	return &OrderRepositoryMysql{
		DB: db,
	}
}

func (r *OrderRepositoryMysql) Create(order *domain.Order) error {
	res := r.DB.Create(&domain.Order{
		ID:     order.ID,
		Date:   order.Date,
		Items:  order.Items,
		Status: order.Status,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *OrderRepositoryMysql) FindAll() ([]*domain.Order, error) {
	var orders []*domain.Order
	res := r.DB.Preload("Items.Product").Find(&orders)

	if res.Error != nil {
		return nil, res.Error
	}

	return orders, nil
}
