package dto

import "github.com/yrnThiago/api-server-go/internal/entity"

type OfferOutputDto struct {
	ID      string
	Price   float64
	Status  entity.OfferStatus
	Product *ProductOutputDto
	Buyer   *UserOutputDto
	Seller  *UserOutputDto
}

func NewOfferOutputDto(offer *entity.Offer) *OfferOutputDto {
	return &OfferOutputDto{
		ID:      offer.ID,
		Price:   offer.Price,
		Status:  offer.Status,
		Product: NewProductOutputDto(offer.Product),
		Buyer:   NewUserOutputDto(offer.Buyer),
		Seller:  NewUserOutputDto(offer.Seller),
	}
}
