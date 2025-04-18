package entity

import (
	"gorm.io/gorm"
)

type OrderRepository interface {
	Add(order *Order) error
	GetMany() ([]*Order, error)
	GetById(id string) (*Order, error)
	UpdateById(order, newOrderBody *Order) (*Order, error)
	DeleteById(id string) error
}

type OrderItems struct {
	OrderID   string  `gorm:"index"                json:"-"`
	ProductID string  `gorm:"index"                json:"-"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Qty       int
}

type Order struct {
	ID     string `gorm:"primaryKey"`
	Status string
	Items  []OrderItems `gorm:"foreignKey:OrderID"`
	gorm.Model
}
