package usecase

import (
	"github.com/google/uuid"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/utils"
	"gorm.io/gorm"
)

var WAITING_PAYMENT = "Aguardando pagamento"

type OrderInputDto struct {
	Items  []OrderItemInputDto `validate:"required,dive"`
	Status string
}

type OrderItemInputDto struct {
	OrderID   string         `gorm:"index"                json:"-"`
	ProductID string         `json:"-"`
	Product   entity.Product `gorm:"foreignKey:ProductID" validate:"required"`
	Qty       int            `json:"qty" validate:"required,min=1"`
}

type OrderOutputDto struct {
	ID     string
	Status string
	Items  []entity.OrderItems
	gorm.Model
}

type OrderUseCase struct {
	orderRepository   entity.OrderRepository
	productRepository entity.ProductRepository
}

func NewOrderUseCase(orderRepository entity.OrderRepository, productRepository entity.ProductRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (u *OrderUseCase) NewOrderItems(items []OrderItemInputDto) []entity.OrderItems {
	var orderItems []entity.OrderItems
	for _, item := range items {
		orderItems = append(orderItems, entity.OrderItems{
			ProductID: item.Product.ID,
			Qty:       item.Qty,
		})
	}

	return orderItems
}

func (u *OrderUseCase) NewOrder(items []OrderItemInputDto, status string) *entity.Order {
	return &entity.Order{
		ID:     uuid.New().String(),
		Status: status,
		Items:  u.NewOrderItems(items),
	}
}

func (u *OrderUseCase) validateOrderItems(items []OrderItemInputDto) error {
	for _, item := range items {
		_, err := u.productRepository.GetById(item.Product.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *OrderUseCase) Add(
	input OrderInputDto,
) (*OrderOutputDto, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, utils.NewErrorInfo("ValidationError", err.Error())
	}

	if err := u.validateOrderItems(input.Items); err != nil {
		return nil, err
	}

	order := u.NewOrder(input.Items, WAITING_PAYMENT)
	err = u.orderRepository.Add(order)
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

func (u *OrderUseCase) GetById(id string) (*entity.Order, error) {
	order, err := u.orderRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *OrderUseCase) UpdateById(
	id string,
	input OrderInputDto,
) (*entity.Order, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, utils.NewErrorInfo("ValidationError", err.Error())
	}

	if err := u.validateOrderItems(input.Items); err != nil {
		return nil, err
	}

	order, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	newOrderBody := u.NewOrder(input.Items, input.Status)
	updatedOrder, err := u.orderRepository.UpdateById(order, newOrderBody)
	if err != nil {
		return nil, err
	}

	return updatedOrder, nil
}

func (u *OrderUseCase) DeleteById(
	id string,
) error {
	err := u.orderRepository.DeleteById(id)
	if err != nil {
		return err
	}

	return err
}
