package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/usecase"
	"github.com/yrnThiago/api-server-go/internal/utils"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "product body missing"})
	}

	validationError := utils.ValidateStruct(input)
	if validationError != nil {
		return utils.NewErrorInfo("ValidationError", validationError.Error())
	}

	output, err := p.ProductUseCase.Add(input)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (p *ProductHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.ProductUseCase.GetMany()
	if err != nil {
		return err
	}

	if output == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "no product was created"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *ProductHandlers) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "product id missing"})
	}

	output, err := p.ProductUseCase.GetById(id)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *ProductHandlers) UpdateById(c *fiber.Ctx) error {
	var input usecase.ProductInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "product body missing"})
	}

	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "product id missing"})
	}

	_, err = p.ProductUseCase.GetById(id)
	if err != nil {
		return err
	}

	new, err := p.ProductUseCase.UpdateById(id, &input)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(new)
}

func (p *ProductHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "product id missing"})
	}
	err := p.ProductUseCase.DeleteById(id)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "product deleted"})
}
