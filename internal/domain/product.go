package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
}

type Product struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Name      string
	Price     float64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(name string, price float64, stock int) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
		Stock: stock,
	}
}
