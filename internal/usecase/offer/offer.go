package usecase

import (
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type OfferUseCase struct {
	OfferRepository IOfferRepository
}

func NewOfferUseCase(offerRepository IOfferRepository) *OfferUseCase {
	return &OfferUseCase{
		OfferRepository: offerRepository,
	}
}

func (u *OfferUseCase) Add(
	input dto.OfferInputDto,
) (*dto.OfferOutputDto, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	offer := entity.NewOffer(
		input.Price,
		entity.PENDING,
		input.ProductID,
		input.SellerID,
		input.BuyerID,
	)
	_, err = u.OfferRepository.Add(offer)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("new offer added")
	return dto.NewOfferOutputDto(offer), nil
}

func (u *OfferUseCase) GetMany() ([]*dto.OfferOutputDto, error) {
	var offersOutputDTO []*dto.OfferOutputDto
	offers, err := u.OfferRepository.GetMany()
	if err != nil {
		return nil, err
	}

	for _, offer := range offers {
		offersOutputDTO = append(offersOutputDTO, dto.NewOfferOutputDto(offer))
	}

	return offersOutputDTO, nil
}

func (u *OfferUseCase) GetById(id string) (*dto.OfferOutputDto, error) {
	offer, err := u.OfferRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewOfferOutputDto(offer), nil
}

func (u *OfferUseCase) GetByUserProductId(userId, productId string) (*dto.OfferOutputDto, error) {
	offer, err := u.OfferRepository.GetByUserProductId(userId, productId)
	if err != nil {
		return nil, err
	}

	return dto.NewOfferOutputDto(offer), nil
}

func (u *OfferUseCase) UpdateById(
	id string,
	input dto.OfferInputDto,
) (*dto.OfferOutputDto, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	offer, err := u.OfferRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	updatedOffer, err := u.OfferRepository.UpdateById(offer)
	if err != nil {
		return nil, err
	}

	config.Logger.Info(
		"offer updated",
		zap.String("id", id),
	)

	return dto.NewOfferOutputDto(updatedOffer), nil
}

func (u *OfferUseCase) DeleteById(
	id string,
) (*dto.OfferOutputDto, error) {
	offer, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	err = u.OfferRepository.DeleteById(id)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("offer deleted",
		zap.String("id", id),
	)
	return offer, err
}
