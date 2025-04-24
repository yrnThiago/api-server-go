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

func (p *PaymentUseCase) GetPaymentMethod(order *entity.Order) entity.PaymentMethod {
	return order.Payment
}

func (p *PaymentUseCase) IsOrderPaymentValid(order *entity.Order) bool {
	return true
}

func (p *PaymentUseCase) GeneratePixCode() string {
	return "PIX_CODE_TEST"
}

func (p *PaymentUseCase) GeneratePayment(order *entity.Order) string {
	if p.GetPaymentMethod(order) == entity.PIX {
		return p.GeneratePixCode()
	}

	return "CARTAO_DE_CREDITO"
}
