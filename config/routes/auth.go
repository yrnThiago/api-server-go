package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/routes"
)

func AuthRouter() *fiber.App {
	authHandlers := handlers.NewAuthHandlers()
	authRouter := routes.NewAuthRouter(authHandlers)

	return authRouter.GetRouter()
}
