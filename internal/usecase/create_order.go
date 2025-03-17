package usecase

import "github.com/yrnThiago/gdlp-go/internal/entity"

type CreateOrderInputDto struct {
	Date  string
	Items []entity.OrderItems
}

type CreateOrderOutputDto struct {
	ID    string
	Date  string
	Items []entity.OrderItems
}

type CreateOrderUseCase struct {
	orderRepository entity.OrderRepository
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (u *CreateOrderUseCase) Execute(
	input CreateOrderInputDto,
) (*CreateOrderOutputDto, error) {
	order := entity.NewOrder(input.Date, input.Items)
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
