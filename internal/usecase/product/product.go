package usecase

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/utils"
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
	ProductRepository IProductRepository
}

func NewProductUseCase(productRepository IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: productRepository,
	}
}

func NewProduct(name string, price float64, stock int) *entity.Product {
	return &entity.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

func (u *ProductUseCase) Add(
	input ProductInputDto,
) (*entity.Product, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, utils.GetValidationError(err.Error())
	}

	product := entity.NewProduct(input.Name, input.Price, input.Stock)
	err = u.ProductRepository.Add(product)
	if err != nil {
		return nil, err
	}

	config.Logger.Info(
		"new product added",
		zap.String("name", input.Name),
	)
	return product, nil
}

func (u *ProductUseCase) GetMany() ([]*entity.Product, error) {
	products, err := u.ProductRepository.GetMany()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *ProductUseCase) GetById(id string) (*entity.Product, error) {
	product, err := u.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductUseCase) UpdateById(
	id string,
	input ProductInputDto,
) (*entity.Product, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, utils.GetValidationError(err.Error())
	}

	product, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	newProductBody := NewProduct(input.Name, input.Price, input.Stock)
	product, err = u.ProductRepository.UpdateById(product, newProductBody)
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
