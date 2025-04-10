package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/cmd/pub"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/usecase"
	"github.com/yrnThiago/api-server-go/internal/utils"
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
		errorInfo := utils.NewErrorInfo(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	output, err := p.OrderUseCase.Add(input)
	if err != nil {
		errorInfo := utils.NewErrorInfo(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	go pub.SendMessage(output)

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (p *OrderHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.OrderUseCase.GetMany()
	if err != nil {
		errorInfo := utils.NewErrorInfo(
			fiber.StatusInternalServerError,
			fiber.ErrInternalServerError.Message,
		)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	c.Set("Content-Type", "application/json")
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
		errorInfo := utils.NewErrorInfo(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OrderHandlers) UpdateById(c *fiber.Ctx) error {
	var input map[string]any
	err := c.BodyParser(&input)
	if err != nil {
		errorInfo := utils.NewErrorInfo(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order id missing"})
	}

	output, err := p.OrderUseCase.GetById(id)
	if err != nil {
		errorInfo := utils.NewErrorInfo(
			fiber.StatusInternalServerError,
			fiber.ErrInternalServerError.Message,
		)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	err = p.OrderUseCase.UpdateById(output, input)
	if err != nil {
		errorInfo := utils.NewErrorInfo(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *OrderHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order id missing"})
	}

	err := p.OrderUseCase.DeleteById(id)
	if err != nil {
		errorInfo := utils.NewErrorInfo(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order deleted"})
}
