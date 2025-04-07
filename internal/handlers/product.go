package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/usecase"
)

type ProductHandlers struct {
	ProductUseCase *usecase.ProductUseCase
}

func NewProductHandlers(
	createProductUseCase *usecase.ProductUseCase,
) *ProductHandlers {
	return &ProductHandlers{
		ProductUseCase: createProductUseCase,
	}
}

func (p *ProductHandlers) Add(c *fiber.Ctx) error {
	var input usecase.ProductInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product body"})
	}

	output, err := p.ProductUseCase.Add(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (p *ProductHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.ProductUseCase.GetMany()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (p *ProductHandlers) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	output, err := p.ProductUseCase.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *ProductHandlers) UpdateById(c *fiber.Ctx) error {
	var input usecase.ProductInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product body"})
	}

	id := c.Params("id")
	_, err = p.ProductUseCase.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "something went wrong"})
	}

	new, err := p.ProductUseCase.UpdateById(id, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(new)
}

func (p *ProductHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	err := p.ProductUseCase.DeleteById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "product deleted"})
}
