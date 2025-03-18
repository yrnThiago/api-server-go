package usecase

import "github.com/yrnThiago/gdlp-go/internal/domain"

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

func (u *ProductUseCase) Create(
	input ProductInputDto,
) (*ProductOutputDto, error) {
	product := domain.NewProduct(input.Name, input.Price, input.Stock)
	err := u.ProductRepository.Create(product)
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
	products, err := u.ProductRepository.FindAll()
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

func (u *ProductUseCase) UpdateById(
	input ProductInputDto,
) (*ProductOutputDto, error) {
	product := domain.NewProduct(input.Name, input.Price, input.Stock)
	err := u.ProductRepository.Create(product)
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
