package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

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
		fmt.Println("Error creating a new tokens", err)
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
