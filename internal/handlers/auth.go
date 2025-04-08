package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type AuthHandler struct{}

func NewAuthHandlers() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	userAuthorization, err := utils.GenerateJWT()
	if err != nil {
		config.Logger.Warn(
			"jwt token not generated",
			zap.Error(err),
		)
	}

	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = userAuthorization
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HTTPOnly = true
	cookie.Path = "/"

	c.Cookie(cookie)

	return c.JSON(TestResponse{"user logged in"})
}
