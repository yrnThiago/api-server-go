package usecase

import (
	"testing"
)

func Test_AddProduct_WithInvalidInput_ShouldReturnValidationError(t *testing.T) {
	productRepository
	productUseCase := NewProductUseCase()
	input := ProductInputDto{
		Name:  "Controle",
		Price: 123.45,
		Stock: 1,
	}

}
