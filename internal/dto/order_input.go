package dto

import "github.com/yrnThiago/api-server-go/internal/entity"

type OrderInputDto struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Items          []OrderItemInputDto  `validate:"required,dive"`
	Status         entity.OrderStatus   `validate:"required,oneof=Aprovado 'Aguardando pagamento' Cancelado"`
	Payment        entity.PaymentMethod `validate:"required,oneof=Pix 'Cartao de credito'"`
	ClientID       string               `json:"client_id" validate:"required"`
}

type OrderItemInputDto struct {
	ProductID string `json:"product_id" validate:"required"`
	Qty       int    `json:"qty" validate:"required,min=1"`
}
