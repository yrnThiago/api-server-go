package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/usecase/user"
)

type UserHandlers struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandlers(
	createUserUseCase *usecase.UserUseCase,
) *UserHandlers {
	return &UserHandlers{
		UserUseCase: createUserUseCase,
	}
}

func (p *UserHandlers) Add(c *fiber.Ctx) error {
	var input dto.UserInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user body missing"})
	}

	_, err = p.UserUseCase.Add(input)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user created successfully"})
}

func (p *UserHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.UserUseCase.GetMany()
	if err != nil {
		return err
	}

	if output == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "no users found"})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *UserHandlers) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user id missing"})
	}

	output, err := p.UserUseCase.GetById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *UserHandlers) UpdateById(c *fiber.Ctx) error {
	var input dto.UserInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user body missing"})
	}

	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user id missing"})
	}

	_, err = p.UserUseCase.UpdateById(id, input)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user updated successfully"})
}

func (p *UserHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user id missing"})
	}

	_, err := p.UserUseCase.DeleteById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user deleted successfully"})
}
