package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type ProductRouter struct {
	Path            string
	ProductHandlers *handlers.ProductHandlers
}

func NewProductRouter(productHandlers *handlers.ProductHandlers) *ProductRouter {
	return &ProductRouter{
		Path:            "/products",
		ProductHandlers: productHandlers,
	}
}

func (p *ProductRouter) GetRouter() chi.Router {
	router := chi.NewRouter()

	router.Post("/", p.ProductHandlers.Add)
	router.Get("/", p.ProductHandlers.GetMany)
	router.Get("/{id}", p.ProductHandlers.GetById)
	router.Put("/{id}", p.ProductHandlers.UpdateById)
	router.Delete("/{id}", p.ProductHandlers.DeleteById)

	return router
}
