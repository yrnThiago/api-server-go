package dto

import "github.com/yrnThiago/api-server-go/internal/entity"

type OfferInputDto struct {
	Price     float64 `validate:"required"`
	ProductID string  `validate:"required" json:"product_id"`
	SellerID  string  `validate:"required" json:"seller_id"`
	BuyerID   string  `validate:"required" json:"buyer_id"`
}

type OfferStatusInputDto struct {
	Status entity.OfferStatus `validate:"required,oneof=Aceita 'Pendente' Recusada"`
}

func NewOfferInputDto(price float64, productId, sellerId, buyerId string) *OfferInputDto {
	return &OfferInputDto{
		Price:     price,
		ProductID: productId,
		SellerID:  sellerId,
		BuyerID:   buyerId,
	}
}
