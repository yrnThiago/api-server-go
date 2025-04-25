package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/factory"
)

func ProductRouter() *fiber.App {
	productFactory := factory.NewProductFactory()

	return productFactory.Router.GetRoutes()
}
