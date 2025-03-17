package usecase

import "github.com/yrnThiago/gdlp-go/internal/domain"


type ListOrdersOutputDto struct {
	ID    string
	Date  string
	Items []domain.OrderItems
}

type ListOrdersUseCase struct {
	orderRepository domain.OrderRepository
}

func NewListOrdersCase(orderRepository domain.OrderRepository) *ListOrdersUseCase {
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
