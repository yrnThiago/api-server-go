package configroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/factory"
)

func OrderRouter() *fiber.App {
	orderFactory := factory.NewOrderFactory()

	return orderFactory.Router.GetRoutes()
}
