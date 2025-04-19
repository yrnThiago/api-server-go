package usecase

import (
	"go.uber.org/zap"

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
) (*ProductOutputDto, error) {
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

	return &ProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}, nil
}

func (u *ProductUseCase) GetMany() ([]*ProductOutputDto, error) {
	var productsDTO []*ProductOutputDto
	products, err := u.ProductRepository.GetMany()
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		productDTO := &ProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}

		productsDTO = append(productsDTO, productDTO)
	}

	return productsDTO, nil
}

func (u *ProductUseCase) GetById(id string) (*ProductOutputDto, error) {
	product, err := u.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &ProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}, nil
}

func (u *ProductUseCase) UpdateById(
	id string,
	input ProductInputDto,
) (*ProductOutputDto, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, utils.GetValidationError(err.Error())
	}

	product, err := u.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock

	updatedProduct, err := u.ProductRepository.UpdateById(product)
	if err != nil {
		return nil, err
	}

	return &ProductOutputDto{
		ID:    updatedProduct.ID,
		Name:  updatedProduct.Name,
		Price: updatedProduct.Price,
		Stock: updatedProduct.Stock,
	}, nil
}

func (u *ProductUseCase) DeleteById(
	id string,
) (*ProductOutputDto, error) {

	product, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	err = u.ProductRepository.DeleteById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
