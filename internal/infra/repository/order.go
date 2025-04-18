package repository

import (
	"errors"
	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/internal/entity"
)

type OrderRepositoryMysql struct {
	DB *gorm.DB
}

func NewOrderRepositoryMysql(db *gorm.DB) *OrderRepositoryMysql {
	return &OrderRepositoryMysql{
		DB: db,
	}
}

func (r *OrderRepositoryMysql) Add(order *entity.Order) error {
	res := r.DB.Create(order).Omit("Items.Product")

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *OrderRepositoryMysql) GetMany() ([]*entity.Order, error) {
	var orders []*entity.Order
	res := r.DB.Preload("Items.Product").Find(&orders)

	if res.Error != nil {
		return nil, res.Error
	}

	return orders, nil
}

func (r *OrderRepositoryMysql) GetById(orderId string) (*entity.Order, error) {
	var order *entity.Order
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
	order, newOrderBody *entity.Order,
) (*entity.Order, error) {
	res := r.DB.Model(&order).Select("status").Omit("ID").Updates(newOrderBody)
	if res.Error != nil {
		return nil, res.Error
	}

	return order, nil
}

func (r *OrderRepositoryMysql) DeleteById(orderId string) error {
	var order *entity.Order
	res := r.DB.Delete(&order, "id = ?", orderId)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
