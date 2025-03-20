package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Add(product *Product) error
	GetMany() ([]*Product, error)
	GetById(id string) (*Product, error)
	UpdateById(id string, newProduct *Product) (*Product, error)
	DeleteById(id string) error
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
