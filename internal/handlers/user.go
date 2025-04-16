package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/usecase"
	"github.com/yrnThiago/api-server-go/internal/utils"
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
	var input usecase.UserInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user body missing"})
	}

	validationError := utils.ValidateStruct(input)
	if validationError != nil {
		return utils.NewErrorInfo("ValidationError", validationError.Error())
	}

	input.Password, err = utils.GenerateHashPassword(input.Password)
	if err != nil {
		return err
	}

	output, err := p.UserUseCase.Add(input)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (p *UserHandlers) GetMany(c *fiber.Ctx) error {
	output, err := p.UserUseCase.GetMany()
	if err != nil {
		return err
	}

	if output == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "no user was created"})
	}

	c.Set("Content-Type", "application/json")
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

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *UserHandlers) GetByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == ":email" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user email missing"})
	}

	output, err := p.UserUseCase.GetByEmail(email)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(output)
}

func (p *UserHandlers) UpdateById(c *fiber.Ctx) error {
	var input usecase.UserInputDto
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user body missing"})
	}

	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user id missing"})
	}

	_, err = p.UserUseCase.GetById(id)
	if err != nil {
		return err
	}

	new, err := p.UserUseCase.UpdateById(id, &input)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(new)
}

func (p *UserHandlers) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == ":id" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user id missing"})
	}
	err := p.UserUseCase.DeleteById(id)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user deleted"})
}
