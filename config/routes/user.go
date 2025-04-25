package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/factory"
)

func UserRouter() *fiber.App {
	userFactory := factory.NewUserFactory()

	return userFactory.Router.GetRoutes()
}
