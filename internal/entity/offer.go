package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OfferStatus string

const (
	ACCEPTED OfferStatus = "Aceita"
	PENDING  OfferStatus = "Pendente"
	DECLINED OfferStatus = "Recusada"
)

type Offer struct {
	ID        string
	Price     float64
	Status    OfferStatus
	ProductID string  `gorm:"index"                json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	SellerID  string  `gorm:"index"                json:"seller_id"`
	Seller    User    `gorm:"foreignKey:SellerID"`
	BuyerID   string  `gorm:"index"                json:"buyer_id"`
	Buyer     User    `gorm:"foreignKey:BuyerID"`
	gorm.Model
}

func NewOffer(price float64, status OfferStatus, productId, sellerId, buyerId string) *Offer {
	return &Offer{
		ID:        uuid.New().String(),
		Price:     price,
		Status:    status,
		ProductID: productId,
		SellerID:  sellerId,
		BuyerID:   buyerId,
	}
}

func (o *Offer) SetAcceptedStatus() {
	o.Status = ACCEPTED
}

func (o *Offer) SetDeclinedStatus() {
	o.Status = DECLINED
}
