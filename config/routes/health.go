package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/routes"
)

func HealthRouter() *fiber.App {
	healthHandlers := handlers.NewHealthHandlers()
	healthRouter := routes.NewHealthRouter(healthHandlers)

	return healthRouter.GetRouter()
}
