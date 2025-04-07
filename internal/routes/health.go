package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type HealthRouter struct {
	Path           string
	HealthHandlers *handlers.HealthHandler
}

func NewHealthRouter(healthHandlers *handlers.HealthHandler) *HealthRouter {
	return &HealthRouter{
		Path:           "/health",
		HealthHandlers: healthHandlers,
	}
}

func (h *HealthRouter) GetRouter() *fiber.App {
	router := fiber.New()
	router.Get("/ping", h.HealthHandlers.Ping)

	return router
}
