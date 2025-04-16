package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/cmd/publisher"
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
		errorInfo := utils.NewErrorInfo(fiber.ErrBadRequest.Message, fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	validationError := utils.ValidateStruct(input)
	if validationError != nil {
		c.Locals(string(keys.ErrorKey), validationError)
		return nil
	}

	output, err := p.OrderUseCase.Add(input)
	if err != nil {
		errorInfo := utils.NewErrorInfo(fiber.ErrBadRequest.Message, fiber.StatusBadRequest, err.Error())
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	go publisher.OrdersPub.Publish(output.ID)

	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusCreated).JSON(output)
	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": fmt.Sprintf("order id: %s created", output.ID)})
}

func (p *OrderHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.OrderUseCase.GetMany()
	if err != nil {
		errorInfo := utils.NewErrorInfo(
			fiber.ErrInternalServerError.Message,
			fiber.StatusInternalServerError,
			fiber.ErrInternalServerError.Message,
		)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	if output == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "no order was created"})
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
		errorInfo := utils.NewErrorInfo(fiber.ErrBadRequest.Message, fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
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
		errorInfo := utils.NewErrorInfo(fiber.ErrBadRequest.Message, fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
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
			fiber.ErrInternalServerError.Message,
			fiber.StatusInternalServerError,
			fiber.ErrInternalServerError.Message,
		)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	err = p.OrderUseCase.UpdateById(output, input)
	if err != nil {
		errorInfo := utils.NewErrorInfo(fiber.ErrBadRequest.Message, fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
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
		errorInfo := utils.NewErrorInfo(fiber.ErrBadRequest.Message, fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order deleted"})
}
