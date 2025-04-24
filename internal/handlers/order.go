package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/infra/nats"
	"github.com/yrnThiago/api-server-go/internal/usecase/order"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order body missing"})
	}

	output, err := p.OrderUseCase.Add(input)
	if err != nil {
		return err
	}

	go nats.OrdersPublisher.Publish(output.ID)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "order created successfully"})
}

func (p *OrderHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.OrderUseCase.GetMany()
	if err != nil {
		return err
	}

	if output == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "no orders found"})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OrderHandlers) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order id missing"})
	}

	output, err := p.OrderUseCase.GetById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OrderHandlers) UpdateById(c *fiber.Ctx) error {
	var input usecase.OrderInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order body missing"})
	}

	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order id missing"})
	}

	_, err = p.OrderUseCase.UpdateById(id, input)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order updated successfully"})
}

func (p *OrderHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order id missing"})
	}

	_, err := p.OrderUseCase.DeleteById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order deleted successfully"})
}
