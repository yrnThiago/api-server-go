package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type ProductRepositoryMysql struct {
	DB *gorm.DB
}

func NewProductRepositoryMysql(db *gorm.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		DB: db,
	}
}

func (r *ProductRepositoryMysql) Add(product *entity.Product) error {
	res := r.DB.Create(product)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *ProductRepositoryMysql) GetMany() ([]*entity.Product, error) {
	var products []*entity.Product
	res := r.DB.Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	return products, nil
}

func (r *ProductRepositoryMysql) GetById(productID string) (*entity.Product, error) {
	var product *entity.Product
	res := r.DB.Limit(1).First(&product, "id = ?", productID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrProductNotFound
		}

		return nil, res.Error
	}

	return product, nil
}

func (r *ProductRepositoryMysql) UpdateById(
	product, newProductBody *entity.Product,
) (*entity.Product, error) {
	r.DB.Model(&product).Omit("ID").Updates(newProductBody)

	return product, nil
}

func (r *ProductRepositoryMysql) DeleteById(productID string) error {
	var product *entity.Product
	res := r.DB.Delete(&product, "id = ?", productID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
