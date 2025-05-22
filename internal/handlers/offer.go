package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/infra/nats"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/offer"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type OfferHandlers struct {
	OfferUseCase *usecase.OfferUseCase
}

func NewOfferHandlers(
	createOfferUseCase *usecase.OfferUseCase,
) *OfferHandlers {
	return &OfferHandlers{
		OfferUseCase: createOfferUseCase,
	}
}

func (p *OfferHandlers) Add(c *fiber.Ctx) error {
	var input dto.OfferInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer body missing"})
	}

	_, err = p.OfferUseCase.Add(input)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "offer created successfully"})
}

func (p *OfferHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.OfferUseCase.GetMany()
	if err != nil {
		return err
	}

	if output == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "no offer found"})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OfferHandlers) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer id missing"})
	}

	output, err := p.OfferUseCase.GetById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OfferHandlers) AcceptById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer id missing"})
	}

	_, err := p.OfferUseCase.GetById(id)
	if err != nil {
		return err
	}

	go nats.OffersPublisher.Publish(nats.OffersAcceptedFilter, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "offer accepted successfully"})
}

func (p *OfferHandlers) DeclineById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer id missing"})
	}

	_, err := p.OfferUseCase.GetById(id)
	if err != nil {
		return err
	}

	go nats.OffersPublisher.Publish(nats.OffersDeclinedFilter, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "offer declined successfully"})
}

func (p *OfferHandlers) AnswerOffer(c *fiber.Ctx) error {
	var input dto.OfferStatusInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer body missing"})
	}

	err = utils.ValidateStruct(input)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer id missing"})
	}

	var offerFilter string

	switch input.Status {
	case entity.ACCEPTED:
		offerFilter = nats.OffersAcceptedFilter
	case entity.PENDING:
		offerFilter = nats.OffersPendingFilter
	case entity.DECLINED:
		offerFilter = nats.OffersDeclinedFilter
	}

	go nats.OffersPublisher.Publish(offerFilter, id)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "offer answered successfully"})
}

func (p *OfferHandlers) UpdateById(c *fiber.Ctx) error {
	var input dto.OfferInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer body missing"})
	}

	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer id missing"})
	}

	_, err = p.OfferUseCase.UpdateById(id, input)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "offer updated successfully"})
}

func (p *OfferHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "offer id missing"})
	}
	_, err := p.OfferUseCase.DeleteById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "offer deleted successfully"})
}
