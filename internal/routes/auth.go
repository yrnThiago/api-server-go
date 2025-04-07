package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type AuthRouter struct {
	Path         string
	AuthHandlers *handlers.AuthHandler
}

func NewAuthRouter(authHandlers *handlers.AuthHandler) *AuthRouter {
	return &AuthRouter{
		Path:         "/auth",
		AuthHandlers: authHandlers,
	}
}

func (h *AuthRouter) GetRouter() *fiber.App {
	router := fiber.New()
	router.Get("/login", h.AuthHandlers.Login)

	return router
}
