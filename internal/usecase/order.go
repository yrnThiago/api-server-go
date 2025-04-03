package usecase

import "github.com/yrnThiago/api-server-go/internal/models"

type OrderInputDto struct {
	Items  []models.OrderItems
	Status string
}

type OrderOutputDto struct {
	ID     string
	Status string
	Items  []models.OrderItems
}

type OrderUseCase struct {
	orderRepository models.OrderRepository
}

func NewOrderUseCase(orderRepository models.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository: orderRepository,
	}
}

func (u *OrderUseCase) Add(
	input OrderInputDto,
) (*OrderOutputDto, error) {
	order := models.NewOrder(input.Items, "Aguardando pagamento")
	err := u.orderRepository.Add(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDto{
		ID:     order.ID,
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
			Items:  order.Items,
			Status: order.Status,
		})
	}

	return ordersOutput, nil
}

func (u *OrderUseCase) GetById(id string) (*models.Order, error) {
	order, err := u.orderRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *OrderUseCase) UpdateById(
	order *models.Order,
	body map[string]any,
) error {
	err := u.orderRepository.UpdateById(order, body)
	if err != nil {
		return err
	}

	return nil
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
