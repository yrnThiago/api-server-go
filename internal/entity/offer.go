package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OfferStatus string

const (
	ACCEPTED OfferStatus = "Aceita"
	PENDING  OfferStatus = "Pendente"
	REFUSED  OfferStatus = "Recusada"
)

type Offer struct {
	ID        string
	Price     float64
	Status    OfferStatus
	ProductID string
	SellerID  string
	BuyerID   string
	gorm.Model
}

func NewOffer(price float64, status OfferStatus, productId, sellerId, buyerId string) *Offer {
	return &Offer{
		ID:        uuid.New().String(),
		Price:     price,
		Status:    status,
		ProductID: productId,
		SellerID:  productId,
		BuyerID:   buyerId,
	}
}
