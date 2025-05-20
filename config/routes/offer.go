package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/factory"
)

func OfferRouter() *fiber.App {
	offerFactory := factory.NewOfferFactory()

	return offerFactory.Router.GetRoutes()
}
