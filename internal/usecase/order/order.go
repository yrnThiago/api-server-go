package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/yrnThiago/api-server-go/internal/entity"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/product"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type OrderInputDto struct {
	Items   []OrderItemInputDto  `validate:"required,dive"`
	Status  entity.OrderStatus   `validate:"required,oneof=Aprovado 'Aguardando pagamento' Cancelado"`
	Payment entity.PaymentMethod `validate:"required,oneof=Pix 'Cartao de credito'"`
}

type OrderItemInputDto struct {
	ProductID string `json:"product_id" validate:"required"`
	Qty       int    `json:"qty" validate:"required,min=1"`
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

func NewOrderProductOutputDto(item entity.OrderItems) *OrderProductOutputDto {
	return &OrderProductOutputDto{
		ID:    item.Product.ID,
		Name:  item.Product.Name,
		Price: item.Product.Price,
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

type OrderOutputDto struct {
	ID        string
	Status    entity.OrderStatus
	Items     []OrderItemsOutputDto
	Payment   entity.PaymentMethod
	CreatedAt time.Time
}

type OrderUseCase struct {
	orderRepository   IOrderRepository
	productRepository usecase.IProductRepository
}

func NewOrderUseCase(orderRepository IOrderRepository, productRepository usecase.IProductRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (u *OrderUseCase) NewOrderItems(items []OrderItemInputDto) []entity.OrderItems {
	var orderItems []entity.OrderItems
	for _, item := range items {
		orderItems = append(orderItems, entity.OrderItems{
			ProductID: item.ProductID,
			Qty:       item.Qty,
		})
	}

	return orderItems
}

func (u *OrderUseCase) NewOrder(items []OrderItemInputDto, status entity.OrderStatus, payment entity.PaymentMethod) *entity.Order {
	return &entity.Order{
		ID:      uuid.New().String(),
		Status:  status,
		Payment: payment,
		Items:   u.NewOrderItems(items),
	}
}

func NewOrderOutputDto(order *entity.Order) *OrderOutputDto {
	return &OrderOutputDto{
		ID:        order.ID,
		Items:     NewOrderItemsOutputDto(order.Items),
		Status:    order.Status,
		Payment:   order.Payment,
		CreatedAt: order.CreatedAt,
	}
}

func (u *OrderUseCase) validateOrderItems(items []OrderItemInputDto) error {
	for _, item := range items {
		_, err := u.productRepository.GetById(item.ProductID)
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
		return nil, entity.GetValidationError(err.Error())
	}

	if err := u.validateOrderItems(input.Items); err != nil {
		return nil, err
	}

	order := u.NewOrder(input.Items, input.Status, input.Payment)
	err = u.orderRepository.Add(order)
	if err != nil {
		return nil, err
	}

	return NewOrderOutputDto(order), nil
}

func (u *OrderUseCase) GetMany() ([]*OrderOutputDto, error) {
	orders, err := u.orderRepository.GetMany()
	if err != nil {
		return nil, err
	}

	var ordersOutput []*OrderOutputDto
	for _, order := range orders {
		ordersOutput = append(ordersOutput, NewOrderOutputDto(order))
	}

	return ordersOutput, nil
}

func (u *OrderUseCase) GetById(id string) (*OrderOutputDto, error) {
	order, err := u.orderRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return NewOrderOutputDto(order), nil
}

func (u *OrderUseCase) UpdateById(
	id string,
	input OrderInputDto,
) (*OrderOutputDto, error) {

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

	return NewOrderOutputDto(updatedOrder), nil
}

func (u *OrderUseCase) DeleteById(
	id string,
) (*OrderOutputDto, error) {
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
