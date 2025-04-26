package dto

import "github.com/yrnThiago/api-server-go/internal/entity"

type ProductOutputDto struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

func NewProductOutputDto(product *entity.Product) *ProductOutputDto {
	return &ProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}
