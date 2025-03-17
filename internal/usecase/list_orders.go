package usecase

import "github.com/yrnThiago/gdlp-go/internal/entity"

type ListOrdersOutputDto struct {
	ID    string
	Date  string
	Items []entity.OrderItems
}

type ListOrdersUseCase struct {
	orderRepository entity.OrderRepository
}

func NewListOrdersCase(orderRepository entity.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		orderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]*ListOrdersOutputDto, error) {
	orders, err := u.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersOutput []*ListOrdersOutputDto
	for _, order := range orders {
		ordersOutput = append(ordersOutput, &ListOrdersOutputDto{
			ID:    order.ID,
			Date:  order.Date,
			Items: order.Items,
		})
	}

	return ordersOutput, nil
}
