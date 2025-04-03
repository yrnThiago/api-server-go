package models

import (
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
	ID    string `gorm:"primaryKey"`
	Name  string
	Price float64
	Stock int
	gorm.Model
}

func NewProduct(name string, price float64, stock int) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
		Stock: stock,
	}
}
