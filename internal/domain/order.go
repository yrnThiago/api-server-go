package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Add(order *Order) error
	GetMany() ([]*Order, error)
	GetById(id string) (*Order, error)
	UpdateById(id string, newOrder *Order) (*Order, error)
	DeleteById(id string) error
}

type OrderItems struct {
	OrderID   string  `gorm:"index"                json:"-"`
	ProductID string  `gorm:"index"                json:"-"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Qty       int
}

// Modelo de Pedido
type Order struct {
	gorm.Model
	ID     string `gorm:"primaryKey"`
	Date   string
	Status string
	Items  []OrderItems `gorm:"foreignKey:OrderID"`
}

func NewOrder(date string, items []OrderItems, status string) *Order {
	return &Order{
		ID:     uuid.New().String(),
		Date:   date,
		Status: status,
		Items:  items,
	}
}
