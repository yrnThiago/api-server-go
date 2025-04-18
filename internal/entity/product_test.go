package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct_return_product_with_id(t *testing.T) {
	productName := "Controle PS4"
	productPrice := 123.45
	productStock := 1

	productModel := NewProduct(productName, productPrice, productStock)

	assert.NotNil(t, productModel)
	assert.NotEmpty(t, productModel.ID)

	_, err := uuid.Parse(productModel.ID)
	assert.NoError(t, err)

	assert.Exactly(t, productName, productModel.Name)
	assert.Exactly(t, productPrice, productModel.Price)
	assert.Exactly(t, productStock, productModel.Stock)
}
