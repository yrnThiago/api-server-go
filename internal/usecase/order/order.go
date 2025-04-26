package usecase

import (
	"github.com/google/uuid"
	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/entity"
	productUseCase "github.com/yrnThiago/api-server-go/internal/usecase/product"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type OrderUseCase struct {
	orderRepository   IOrderRepository
	productRepository productUseCase.IProductRepository
}

func NewOrderUseCase(orderRepository IOrderRepository, productRepository productUseCase.IProductRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (u *OrderUseCase) NewOrderItems(items []dto.OrderItemInputDto) []entity.OrderItems {
	var orderItems []entity.OrderItems
	for _, item := range items {
		orderItems = append(orderItems, entity.OrderItems{
			ProductID: item.ProductID,
			Qty:       item.Qty,
		})
	}

	return orderItems
}

func (u *OrderUseCase) NewOrder(items []dto.OrderItemInputDto, status entity.OrderStatus, payment entity.PaymentMethod, clientID string) *entity.Order {
	return &entity.Order{
		ID:       uuid.New().String(),
		Status:   status,
		ClientID: clientID,
		Payment:  payment,
		Items:    u.NewOrderItems(items),
	}
}

func (u *OrderUseCase) validateOrderItems(items []dto.OrderItemInputDto) error {
	for _, item := range items {
		_, err := u.productRepository.GetById(item.ProductID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *OrderUseCase) Add(
	input dto.OrderInputDto,
) (*dto.OrderOutputDto, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, entity.GetValidationError(err.Error())
	}

	if err := u.validateOrderItems(input.Items); err != nil {
		return nil, err
	}

	order := u.NewOrder(input.Items, input.Status, input.Payment, input.ClientID)
	err = u.orderRepository.Add(order)
	if err != nil {
		return nil, err
	}

	return dto.NewOrderOutputDto(order), nil
}

func (u *OrderUseCase) GetMany() ([]*dto.OrderOutputDto, error) {
	orders, err := u.orderRepository.GetMany()
	if err != nil {
		return nil, err
	}

	var ordersOutput []*dto.OrderOutputDto
	for _, order := range orders {
		ordersOutput = append(ordersOutput, dto.NewOrderOutputDto(order))
	}

	return ordersOutput, nil
}

func (u *OrderUseCase) GetById(id string) (*dto.OrderOutputDto, error) {
	order, err := u.orderRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewOrderOutputDto(order), nil
}

func (u *OrderUseCase) UpdateById(
	id string,
	input dto.OrderInputDto,
) (*dto.OrderOutputDto, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, entity.GetValidationError(err.Error())
	}

	if err := u.validateOrderItems(input.Items); err != nil {
		return nil, err
	}

	order, err := u.orderRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	order.Status = input.Status

	updatedOrder, err := u.orderRepository.UpdateById(order)
	if err != nil {
		return nil, err
	}

	return dto.NewOrderOutputDto(updatedOrder), nil
}

func (u *OrderUseCase) DeleteById(
	id string,
) (*dto.OrderOutputDto, error) {
	order, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	err = u.orderRepository.DeleteById(id)
	if err != nil {
		return nil, err
	}

	return order, err
}
