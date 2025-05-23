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

	Pix        PaymentMethod = "Pix"
	CreditCard PaymentMethod = "Cartao de credito"
)

type OrderItems struct {
	OrderID   string  `gorm:"index"                json:"-"`
	ProductID string  `gorm:"index"                json:"-"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Qty       int
}

type Order struct {
	ID             string `gorm:"primaryKey"`
	IdempotencyKey string `gorm:"unique" json:"idempotency_key"`
	Status         OrderStatus
	Items          []OrderItems `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	ClientID       string       `gorm:"index" json:"-"`
	Client         User         `gorm:"foreignKey:ClientID"`
	Payment        PaymentMethod
	gorm.Model
}

func (o *Order) SetStatus(status OrderStatus) {
	o.Status = status
}
