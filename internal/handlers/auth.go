package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/usecase/auth"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type AuthHandler struct {
	AuthUseCase usecase.IAuthUseCase
}

func NewAuthHandlers(authUseCase usecase.IAuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: authUseCase,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input usecase.AuthInputDto
	c.BodyParser(&input)

	token, output, err := h.AuthUseCase.Login(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong credentials"})
	}

	c.Locals(utils.UserIdKeyCtx, output.ID)

	config.Logger.Info(
		"user logged in",
		zap.String("user id", c.Locals(utils.UserIdKeyCtx).(string)),
	)

	utils.SetBearerCookie(c, token)
	return c.JSON(fiber.Map{"message": "user logged in"})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	utils.ClearBearerCookie(c)

	return c.JSON(fiber.Map{"message": "user logout"})
}
