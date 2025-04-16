package usecase

import (
	"github.com/google/uuid"
	"github.com/yrnThiago/api-server-go/internal/models"
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
	Product   models.Product `gorm:"foreignKey:ProductID" validate:"required"`
	Qty       int            `json:"qty" validate:"required,min=1"`
}

type OrderOutputDto struct {
	ID     string
	Status string
	Items  []models.OrderItems
	gorm.Model
}

type OrderUseCase struct {
	orderRepository   models.OrderRepository
	productRepository models.ProductRepository
}

func NewOrderUseCase(orderRepository models.OrderRepository, productRepository models.ProductRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (u *OrderUseCase) NewOrderItems(items []OrderItemInputDto) []models.OrderItems {
	var orderItems []models.OrderItems
	for _, item := range items {
		orderItems = append(orderItems, models.OrderItems{
			ProductID: item.Product.ID,
			Qty:       item.Qty,
		})
	}

	return orderItems
}

func (u *OrderUseCase) NewOrder(items []OrderItemInputDto, status string) *models.Order {
	return &models.Order{
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

	if err := u.validateOrderItems(input.Items); err != nil {
		return nil, err
	}

	order := u.NewOrder(input.Items, WAITING_PAYMENT)
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
