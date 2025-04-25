package entity

import (
	"gorm.io/gorm"
)

type OrderStatus string
type PaymentMethod string

const (
	Paid     OrderStatus = "Aprovado"
	Pending  OrderStatus = "Aguardando pagamento"
	Canceled OrderStatus = "Cancelado"

	PIX         PaymentMethod = "Pix"
	CREDIT_CARD PaymentMethod = "Cartao de credito"
)

type OrderItems struct {
	OrderID   string  `gorm:"index"                json:"-"`
	ProductID string  `gorm:"index"                json:"-"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Qty       int
}

type Order struct {
	ID      string `gorm:"primaryKey"`
	Status  OrderStatus
	Items   []OrderItems `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Payment PaymentMethod
	gorm.Model
}

func (o *Order) SetStatus(status OrderStatus) {
	o.Status = status
}
