package handlers

import (
	"time"

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

	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = utils.BEARER_KEY + token
	cookie.Expires = time.Now().Add(config.Env.COOKIE_EXPIRES_AT)
	cookie.Secure = false
	cookie.HTTPOnly = true
	cookie.Path = "/"

	utils.SetCookie(c, cookie)
	return c.JSON(fiber.Map{"message": "user logged in"})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = "deleted"
	cookie.Expires = time.Now().Add(-3 * time.Second)
	cookie.HTTPOnly = true

	utils.SetCookie(c, cookie)
	return c.JSON(fiber.Map{"message": "user logout"})
}
