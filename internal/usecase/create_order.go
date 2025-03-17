package usecase

import "github.com/yrnThiago/gdlp-go/internal/domain"

type CreateOrderInputDto struct {
	Date  string
	Items []domain.OrderItems
}

type CreateOrderOutputDto struct {
	ID    string
	Date  string
	Items []domain.OrderItems
}

type CreateOrderUseCase struct {
	orderRepository domain.OrderRepository
}

func NewCreateOrderUseCase(orderRepository domain.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (u *CreateOrderUseCase) Execute(
	input CreateOrderInputDto,
) (*CreateOrderOutputDto, error) {
	order := domain.NewOrder(input.Date, input.Items)
	err := u.orderRepository.Create(order)
	if err != nil {
		return nil, err
	}

	return &CreateOrderOutputDto{
		ID:    order.ID,
		Date:  order.Date,
		Items: order.Items,
	}, err
}
