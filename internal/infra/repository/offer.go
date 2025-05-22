package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/internal/entity"
)

type OfferRepositoryMysql struct {
	DB *gorm.DB
}

func NewOfferRepositoryMysql(db *gorm.DB) *OfferRepositoryMysql {
	return &OfferRepositoryMysql{
		DB: db,
	}
}

func (r *OfferRepositoryMysql) Add(offer *entity.Offer) (*entity.Offer, error) {
	res := r.DB.Create(offer)

	if res.Error != nil {
		return nil, res.Error
	}

	return offer, nil
}

func (r *OfferRepositoryMysql) GetMany() ([]*entity.Offer, error) {
	var offers []*entity.Offer
	res := r.DB.Preload("Product").Preload("Buyer").Preload("Seller").Find(&offers)

	if res.Error != nil {
		return nil, res.Error
	}

	return offers, nil
}

func (r *OfferRepositoryMysql) GetById(offerID string) (*entity.Offer, error) {
	var offer *entity.Offer
	res := r.DB.Limit(1).
		Preload("Product").
		Preload("Buyer").
		Preload("Seller").
		First(&offer, "id = ?", offerID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, entity.GetNotFoundError()
		}

		return nil, res.Error
	}

	return offer, nil
}

func (r *OfferRepositoryMysql) GetByUserProductId(userId, productId string) (*entity.Offer, error) {
	var offer *entity.Offer
	res := r.DB.Limit(1).
		Where("Status = ? AND product_id = ? AND buyer_id = ?", entity.ACCEPTED, productId, userId).
		Preload("Product").
		Preload("Buyer").
		Preload("Seller").
		First(&offer)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, entity.GetNotFoundError()
		}

		return nil, res.Error
	}

	return offer, nil
}

func (r *OfferRepositoryMysql) UpdateById(offer *entity.Offer) (*entity.Offer, error) {
	res := r.DB.Save(offer)
	if res.Error != nil {
		return nil, res.Error
	}

	return offer, nil
}

func (r *OfferRepositoryMysql) DeleteById(offerID string) error {
	var offer *entity.Offer
	res := r.DB.Delete(&offer, "id = ?", offerID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
