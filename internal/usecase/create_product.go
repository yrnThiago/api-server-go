package usecase

import "github.com/yrnThiago/gdlp-go/internal/domain"

type CreateProductInputDto struct {
	Name  string
	Price float64
	Stock int
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

type CreateProductUseCase struct {
	ProductRepository domain.ProductRepository
}

func NewCreateProductUseCase(productRepository domain.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *CreateProductUseCase) Execute(
	input CreateProductInputDto,
) (*CreateProductOutputDto, error) {
	product := domain.NewProduct(input.Name, input.Price, input.Stock)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}, err
}
