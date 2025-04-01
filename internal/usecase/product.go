package usecase

import (
	"gorm.io/gorm"

	"github.com/yrnThiago/gdlp-go/internal/domain"
)

type ProductInputDto struct {
	Name  string
	Price float64
	Stock int
}

type ProductOutputDto struct {
	gorm.Model
	ID    string
	Name  string
	Price float64
	Stock int
}

type ProductUseCase struct {
	ProductRepository domain.ProductRepository
}

func NewProductUseCase(productRepository domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *ProductUseCase) Add(
	input ProductInputDto,
) (*domain.Product, error) {
	product := domain.NewProduct(input.Name, input.Price, input.Stock)
	err := u.ProductRepository.Add(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductUseCase) GetMany() ([]*domain.Product, error) {
	products, err := u.ProductRepository.GetMany()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *ProductUseCase) GetById(id string) (*domain.Product, error) {
	product, err := u.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductUseCase) UpdateById(
	productId string,
	input *ProductInputDto,
) (*domain.Product, error) {
	newProduct := domain.NewProduct(input.Name, input.Price, input.Stock)
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
