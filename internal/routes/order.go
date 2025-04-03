package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type OrderRouter struct {
	Path          string
	OrderHandlers *handlers.OrderHandlers
}

func NewOrderRouter(orderHandlers *handlers.OrderHandlers) *OrderRouter {
	return &OrderRouter{
		Path:          "/orders",
		OrderHandlers: orderHandlers,
	}
}

func (o *OrderRouter) GetRouter() chi.Router {
	router := chi.NewRouter()

	router.Post("/checkout", o.OrderHandlers.Add)
	router.Get("/", o.OrderHandlers.GetMany)
	router.Get("/{id}", o.OrderHandlers.GetById)
	router.Put("/{id}", o.OrderHandlers.UpdateById)
	router.Delete("/{id}", o.OrderHandlers.DeleteById)

	return router
}
