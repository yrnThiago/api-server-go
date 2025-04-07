package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/cmd/pub"
	"github.com/yrnThiago/api-server-go/internal/usecase"
)

type OrderHandlers struct {
	OrderUseCase *usecase.OrderUseCase
}

func NewOrderHandlers(
	createOrderUseCase *usecase.OrderUseCase,
) *OrderHandlers {
	return &OrderHandlers{
		OrderUseCase: createOrderUseCase,
	}
}

func (p *OrderHandlers) Add(c *fiber.Ctx) error {
	var input usecase.OrderInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid order body"})
	}

	output, err := p.OrderUseCase.Add(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid order body"})
	}

	go pub.SendMessage(output)

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "new order created"})
}

func (p *OrderHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.OrderUseCase.GetMany()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OrderHandlers) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	output, err := p.OrderUseCase.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OrderHandlers) UpdateById(c *fiber.Ctx) error {
	var input map[string]any
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid order body"})
	}

	id := c.Params("id")
	output, err := p.OrderUseCase.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}

	err = p.OrderUseCase.UpdateById(output, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OrderHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	err := p.OrderUseCase.DeleteById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order deleted"})
}
