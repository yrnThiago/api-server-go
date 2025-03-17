package repository

import (
	"gorm.io/gorm"

	"github.com/yrnThiago/gdlp-go/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *gorm.DB
}

func NewProductRepositoryMysql(db *gorm.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		DB: db,
	}
}

func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	res := r.DB.Create(&entity.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	var products []*entity.Product
	res := r.DB.Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	return products, nil
}
