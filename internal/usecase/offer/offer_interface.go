package usecase

import "github.com/yrnThiago/api-server-go/internal/entity"

type IOfferRepository interface {
	Add(offer *entity.Offer) (*entity.Offer, error)
	GetMany() ([]*entity.Offer, error)
	GetById(id string) (*entity.Offer, error)
	GetByUserProductId(userId, productId string) (*entity.Offer, error)
	UpdateById(offer *entity.Offer) (*entity.Offer, error)
	DeleteById(id string) error
}
