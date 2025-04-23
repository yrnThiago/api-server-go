package entity

import (
	"gorm.io/gorm"
)

type OrderStatus string

const (
	Paid     OrderStatus = "Aprovado"
	Pending  OrderStatus = "Aguardando pagamento"
	Canceled OrderStatus = "Cancelado"
)

type OrderItems struct {
	OrderID   string  `gorm:"index"                json:"-"`
	ProductID string  `gorm:"index"                json:"-"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Qty       int
}

type Order struct {
	ID     string       `gorm:"primaryKey"`
	Status OrderStatus  `gorm:"type:enum('Aprovado', 'Aguardando pagamento', 'Cancelado')"`
	Items  []OrderItems `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	gorm.Model
}
