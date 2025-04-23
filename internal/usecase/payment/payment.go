package payment

import (
	"github.com/yrnThiago/api-server-go/internal/entity"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/order"
)

type PaymentUseCase struct {
	OrderRepository usecase.IOrderRepository
}

func NewPaymentUseCase(orderRepository usecase.IOrderRepository) *PaymentUseCase {
	return &PaymentUseCase{
		OrderRepository: orderRepository,
	}
}

func (p *PaymentUseCase) IsOrderPaymentValid(order *entity.Order) bool {
	return false
}

func (p *PaymentUseCase) GeneratePixCode() string {
	return "12345"
}
