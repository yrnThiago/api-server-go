package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/internal/models"
)

type OrderRepositoryMysql struct {
	DB *gorm.DB
}

func NewOrderRepositoryMysql(db *gorm.DB) *OrderRepositoryMysql {
	return &OrderRepositoryMysql{
		DB: db,
	}
}

func (r *OrderRepositoryMysql) Add(order *models.Order) error {
	res := r.DB.Create(order)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *OrderRepositoryMysql) GetMany() ([]*models.Order, error) {
	var orders []*models.Order
	res := r.DB.Preload("Items.Product").Find(&orders)

	if res.Error != nil {
		return nil, res.Error
	}

	return orders, nil
}

func (r *OrderRepositoryMysql) GetById(orderId string) (*models.Order, error) {
	var order *models.Order
	res := r.DB.Preload("Items.Product").First(&order, "id = ?", orderId)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		return nil, res.Error
	}

	return order, nil
}

func (r *OrderRepositoryMysql) UpdateById(
	order *models.Order,
	body map[string]any,
) error {
	fmt.Println(body)
	r.DB.Model(&order).Omit("Items").Updates(body)

	return nil
}

func (r *OrderRepositoryMysql) DeleteById(orderId string) error {
	var order *models.Order
	res := r.DB.Delete(&order, "id = ?", orderId)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
