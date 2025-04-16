package usecase

import (
	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/models"
)

type ProductInputDto struct {
	Name  string  `validate:"required"`
	Price float64 `validate:"required,gt=0"`
	Stock int     `validate:"required"`
}

type ProductOutputDto struct {
	ID    string
	Name  string
	Price float64
	Stock int
	gorm.Model
}

type ProductUseCase struct {
	ProductRepository models.ProductRepository
}

func NewProductUseCase(productRepository models.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *ProductUseCase) Add(
	input ProductInputDto,
) (*models.Product, error) {
	product := models.NewProduct(input.Name, input.Price, input.Stock)
	err := u.ProductRepository.Add(product)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("adding new product")
	return product, nil
}

func (u *ProductUseCase) GetMany() ([]*models.Product, error) {
	products, err := u.ProductRepository.GetMany()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *ProductUseCase) GetById(id string) (*models.Product, error) {
	product, err := u.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductUseCase) UpdateById(
	productId string,
	input *ProductInputDto,
) (*models.Product, error) {
	newProduct := models.NewProduct(input.Name, input.Price, input.Stock)
	product, err := u.ProductRepository.UpdateById(productId, newProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductUseCase) DeleteById(
	productId string,
) error {
	err := u.ProductRepository.DeleteById(productId)
	if err != nil {
		return err
	}

	return err
}
