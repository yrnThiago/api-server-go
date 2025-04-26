package dto

import (
	"github.com/yrnThiago/api-server-go/internal/entity"
)

type UserOutputDto struct {
	ID     string
	Email  string
	Orders []*UserOrderOutputDto
}

type UserOrderOutputDto struct {
	ID      string
	Status  entity.OrderStatus
	Items   []OrderItemsOutputDto
	Payment entity.PaymentMethod
}

func NewUserOutputDto(user *entity.User) *UserOutputDto {
	return &UserOutputDto{
		ID:     user.ID,
		Email:  user.Email,
		Orders: NewOrdersOutputDto(user.Orders),
	}
}

func NewUserOrderOutputDto(order entity.Order) *UserOrderOutputDto {
	return &UserOrderOutputDto{
		ID:      order.ID,
		Status:  order.Status,
		Items:   NewOrderItemsOutputDto(order.Items),
		Payment: order.Payment,
	}
}

func NewOrdersOutputDto(orders []entity.Order) []*UserOrderOutputDto {
	var ordersOutput []*UserOrderOutputDto

	for _, order := range orders {
		ordersOutput = append(ordersOutput, NewUserOrderOutputDto(order))
	}

	return ordersOutput
}
