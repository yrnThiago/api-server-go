package usecase

import "github.com/yrnThiago/gdlp-go/internal/domain"

type OrderInputDto struct {
	Date  string
	Items []domain.OrderItems
}

type OrderOutputDto struct {
	ID     string
	Date   string
	Status string
	Items  []domain.OrderItems
}

type OrderUseCase struct {
	orderRepository domain.OrderRepository
}

func NewOrderUseCase(orderRepository domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository: orderRepository,
	}
}

func (u *OrderUseCase) Create(
	input OrderInputDto,
) (*OrderOutputDto, error) {
	order := domain.NewOrder(input.Date, input.Items, "Aguardando pagamento")
	err := u.orderRepository.Create(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDto{
		ID:     order.ID,
		Date:   order.Date,
		Items:  order.Items,
		Status: order.Status,
	}, err
}

func (u *OrderUseCase) GetMany() ([]*OrderOutputDto, error) {
	orders, err := u.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersOutput []*OrderOutputDto
	for _, order := range orders {
		ordersOutput = append(ordersOutput, &OrderOutputDto{
			ID:     order.ID,
			Date:   order.Date,
			Items:  order.Items,
			Status: order.Status,
		})
	}

	return ordersOutput, nil
}
