package dto

import (
	"time"

	"github.com/yrnThiago/api-server-go/internal/entity"
)

type OrderOutputDto struct {
	ID        string
	Status    entity.OrderStatus
	Items     []OrderItemsOutputDto
	Payment   entity.PaymentMethod
	Client    *UserOutputDto
	CreatedAt time.Time
}

type OrderItemsOutputDto struct {
	Product *OrderProductOutputDto
	Qty     int
}

type OrderProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

func NewOrderOutputDto(order *entity.Order) *OrderOutputDto {
	return &OrderOutputDto{
		ID:        order.ID,
		Items:     NewOrderItemsOutputDto(order.Items),
		Status:    order.Status,
		Client:    NewUserOutputDto(&order.Client),
		Payment:   order.Payment,
		CreatedAt: order.CreatedAt,
	}
}

func NewOrderUserOutputDto(order *entity.Order) *OrderOutputDto {
	return &OrderOutputDto{
		ID:        order.ID,
		Items:     NewOrderItemsOutputDto(order.Items),
		Status:    order.Status,
		Payment:   order.Payment,
		CreatedAt: order.CreatedAt,
	}
}

func NewOrderItemsOutputDto(items []entity.OrderItems) []OrderItemsOutputDto {
	var itemsOutput []OrderItemsOutputDto

	for _, item := range items {
		itemsOutput = append(itemsOutput, OrderItemsOutputDto{
			Product: NewOrderProductOutputDto(item),
			Qty:     item.Qty},
		)
	}

	return itemsOutput
}

func NewOrderProductOutputDto(item entity.OrderItems) *OrderProductOutputDto {
	return &OrderProductOutputDto{
		ID:    item.Product.ID,
		Name:  item.Product.Name,
		Price: item.Product.Price,
	}
}
