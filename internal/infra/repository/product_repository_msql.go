package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/internal/models"
)

type ProductRepositoryMysql struct {
	DB *gorm.DB
}

func NewProductRepositoryMysql(db *gorm.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		DB: db,
	}
}

func (r *ProductRepositoryMysql) Add(product *models.Product) error {
	res := r.DB.Create(product)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *ProductRepositoryMysql) GetMany() ([]*models.Product, error) {
	var products []*models.Product
	res := r.DB.Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	return products, nil
}

func (r *ProductRepositoryMysql) GetById(productID string) (*models.Product, error) {
	var product *models.Product
	res := r.DB.Limit(1).First(&product, "id = ?", productID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		return nil, res.Error
	}

	return product, nil
}

func (r *ProductRepositoryMysql) UpdateById(
	productID string,
	newProduct *models.Product,
) (*models.Product, error) {
	product, err := r.GetById(productID)
	if err != nil {
		return nil, err
	}

	r.DB.Model(&product).Omit("ID").Updates(newProduct)

	return product, nil
}

func (r *ProductRepositoryMysql) DeleteById(productID string) error {
	var product *models.Product
	res := r.DB.Delete(&product, "id = ?", productID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
