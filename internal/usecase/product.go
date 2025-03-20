package usecase

import (
	"github.com/yrnThiago/gdlp-go/internal/domain"
)

type ProductInputDto struct {
	Name  string
	Price float64
	Stock int
}

type ProductOutputDto struct {
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
) (*ProductOutputDto, error) {
	product := domain.NewProduct(input.Name, input.Price, input.Stock)
	err := u.ProductRepository.Add(product)
	if err != nil {
		return nil, err
	}

	return &ProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}, err
}

func (u *ProductUseCase) GetMany() ([]*ProductOutputDto, error) {
	products, err := u.ProductRepository.GetMany()
	if err != nil {
		return nil, err
	}

	var productsOutput []*ProductOutputDto
	for _, product := range products {
		productsOutput = append(productsOutput, &ProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		})
	}

	return productsOutput, nil
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
	}, err
}

func (u *ProductUseCase) UpdateById(
	productId string,
	input *ProductInputDto,
) (*ProductOutputDto, error) {
	newProduct := domain.NewProduct(input.Name, input.Price, input.Stock)
	_, err := u.ProductRepository.UpdateById(productId, newProduct)
	if err != nil {
		return nil, err
	}

	return &ProductOutputDto{
		ID:    newProduct.ID,
		Name:  newProduct.Name,
		Price: newProduct.Price,
		Stock: newProduct.Stock,
	}, err
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
