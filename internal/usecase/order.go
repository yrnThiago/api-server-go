package usecase

import "github.com/yrnThiago/gdlp-go/internal/domain"

type OrderInputDto struct {
	Date   string
	Items  []domain.OrderItems
	Status string
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

func (u *OrderUseCase) Add(
	input OrderInputDto,
) (*OrderOutputDto, error) {
	order := domain.NewOrder(input.Date, input.Items, "Aguardando pagamento")
	err := u.orderRepository.Add(order)
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
	orders, err := u.orderRepository.GetMany()
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

func (u *OrderUseCase) GetById(id string) (*OrderOutputDto, error) {
	order, err := u.orderRepository.GetById(id)
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

func (u *OrderUseCase) UpdateById(
	orderId string,
	input *OrderInputDto,
) (*OrderOutputDto, error) {
	// u dont need to create a new order, fix later
	newOrder := domain.NewOrder(input.Date, input.Items, input.Status)
	_, err := u.orderRepository.UpdateById(orderId, newOrder)
	if err != nil {
		return nil, err
	}
	return &OrderOutputDto{
		ID:     newOrder.ID,
		Date:   newOrder.Date,
		Items:  newOrder.Items,
		Status: newOrder.Status,
	}, err

}

func (u *OrderUseCase) DeleteById(
	productId string,
) error {
	err := u.orderRepository.DeleteById(productId)
	if err != nil {
		return err
	}

	return err
}
