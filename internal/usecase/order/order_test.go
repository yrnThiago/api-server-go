package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yrnThiago/api-server-go/internal/entity"
)

func TestNewOrder_return_order_with_id(t *testing.T) {
	orderStatus := WAITING_PAYMENT
	orderItems := []entity.OrderItems{
		{
			Product: *entity.NewProduct("Controle PS4", 123.45, 1),
			Qty:     1,
		},
	}

	orderModel := NewOrder(orderItems, orderStatus)

	assert.NotNil(t, orderModel)
	assert.NotEmpty(t, orderModel.ID)

	_, err := uuid.Parse(orderModel.ID)
	assert.NoError(t, err)

	assert.Exactly(t, orderStatus, orderModel.Status)
	assert.GreaterOrEqual(t, len(orderModel.Items), 1)
	assert.Exactly(t, orderItems[0].Qty, orderModel.Items[0].Qty)
	assert.Equal(t, orderItems[0].Product.Name, orderModel.Items[0].Product.Name)
	assert.Equal(t, orderItems[0].Product.Price, orderModel.Items[0].Product.Price)
}
