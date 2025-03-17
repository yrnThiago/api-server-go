package repository

import (
	"github.com/yrnThiago/gdlp-go/internal/domain"
	"gorm.io/gorm"
)

type ProductRepositoryMysql struct {
	DB *gorm.DB
}

func NewProductRepositoryMysql(db *gorm.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		DB: db,
	}
}

func (r *ProductRepositoryMysql) Create(product *domain.Product) error {
	res := r.DB.Create(&domain.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product
	res := r.DB.Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	return products, nil
}
