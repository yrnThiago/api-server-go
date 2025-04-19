package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type UserRouter struct {
	Path         string
	UserHandlers *handlers.UserHandlers
}

func NewUserRouter(userHandlers *handlers.UserHandlers) *UserRouter {
	return &UserRouter{
		Path:         "/users",
		UserHandlers: userHandlers,
	}
}

func (p *UserRouter) GetRouter() *fiber.App {
	router := fiber.New()

	router.Post("/", p.UserHandlers.Add)
	router.Get("/", p.UserHandlers.GetMany)
	router.Get("/:id", p.UserHandlers.GetById)
	router.Put("/:id", p.UserHandlers.UpdateById)
	router.Delete("/:id", p.UserHandlers.DeleteById)

	return router
}
