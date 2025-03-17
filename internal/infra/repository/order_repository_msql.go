package repository

import (
	"gorm.io/gorm"

	"github.com/yrnThiago/gdlp-go/internal/entity"
)

type OrderRepositoryMysql struct {
	DB *gorm.DB
}

func NewOrderRepositoryMysql(db *gorm.DB) *OrderRepositoryMysql {
	return &OrderRepositoryMysql{
		DB: db,
	}
}

func (r *OrderRepositoryMysql) Create(order *entity.Order) error {
	res := r.DB.Create(&entity.Order{
		ID:    order.ID,
		Date:  order.Date,
		Items: order.Items,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *OrderRepositoryMysql) FindAll() ([]*entity.Order, error) {
	var orders []*entity.Order
	res := r.DB.Preload("Items.Product").Find(&orders)

	if res.Error != nil {
		return nil, res.Error
	}

	return orders, nil
}
