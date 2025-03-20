package repository

import (
	"errors"

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

func (r *OrderRepositoryMysql) Add(order *domain.Order) error {
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

func (r *OrderRepositoryMysql) GetMany() ([]*domain.Order, error) {
	var orders []*domain.Order
	res := r.DB.Preload("Items.Product").Find(&orders)

	if res.Error != nil {
		return nil, res.Error
	}

	return orders, nil
}

func (r *OrderRepositoryMysql) GetById(orderId string) (*domain.Order, error) {
	var order *domain.Order
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
	orderId string,
	newOrder *domain.Order,
) (*domain.Order, error) {
	order, err := r.GetById(orderId)
	if err != nil {
		return nil, err
	}

	if newOrder.Date != "" {
		order.Date = newOrder.Date
	}
	if newOrder.Status != "" {
		order.Status = newOrder.Status
	}

	if newOrder.Items != nil {
		order.Items = newOrder.Items
	}

	r.DB.Save(&order)

	return order, nil
}

func (r *OrderRepositoryMysql) DeleteById(orderId string) error {
	var order *domain.Order
	res := r.DB.Delete(&order, "id = ?", orderId)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
