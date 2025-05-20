package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type OfferRouter struct {
	Path          string
	OfferHandlers *handlers.OfferHandlers
}

func NewOfferRouter(offerHandlers *handlers.OfferHandlers) *OfferRouter {
	return &OfferRouter{
		Path:          "/users",
		OfferHandlers: offerHandlers,
	}
}

func (p *OfferRouter) GetRoutes() *fiber.App {
	router := fiber.New()

	router.Post("/", p.OfferHandlers.Add)
	router.Get("/", p.OfferHandlers.GetMany)
	router.Get("/:id", p.OfferHandlers.GetById)
	router.Put("/:id", p.OfferHandlers.UpdateById)
	router.Delete("/:id", p.OfferHandlers.DeleteById)

	return router
}
