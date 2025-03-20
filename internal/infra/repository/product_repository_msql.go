package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/yrnThiago/gdlp-go/internal/domain"
)

type ProductRepositoryMysql struct {
	DB *gorm.DB
}

func NewProductRepositoryMysql(db *gorm.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		DB: db,
	}
}

func (r *ProductRepositoryMysql) Add(product *domain.Product) error {
	res := r.DB.Create(&domain.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *ProductRepositoryMysql) GetMany() ([]*domain.Product, error) {
	var products []*domain.Product
	res := r.DB.Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	return products, nil
}

func (r *ProductRepositoryMysql) GetById(productID string) (*domain.Product, error) {
	var product *domain.Product
	res := r.DB.First(&product, "id = ?", productID)
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
	newProduct *domain.Product,
) (*domain.Product, error) {
	product, err := r.GetById(productID)
	if err != nil {
		return nil, err
	}

	r.DB.Model(&product).Omit("ID").Updates(newProduct)

	return product, nil
}

func (r *ProductRepositoryMysql) DeleteById(productID string) error {
	var product *domain.Product
	res := r.DB.Delete(&product, "id = ?", productID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
