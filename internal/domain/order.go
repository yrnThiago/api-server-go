package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *Order) error
	FindAll() ([]*Order, error)
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
	ID    string `gorm:"primaryKey"`
	Date  string
	Items []OrderItems `gorm:"foreignKey:OrderID"`
}

func NewOrder(date string, items []OrderItems) *Order {
	return &Order{
		ID:    uuid.New().String(),
		Date:  date,
		Items: items,
	}
}
