package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type ProtectedRouter struct {
	Path              string
	ProtectedHandlers *handlers.ProtectedHandler
}

func NewProtectedRouter(protectedHandlers *handlers.ProtectedHandler) *ProtectedRouter {
	return &ProtectedRouter{
		Path:              "/protected",
		ProtectedHandlers: protectedHandlers,
	}
}

func (h *ProtectedRouter) GetRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/", h.ProtectedHandlers.TestCtx)

	return router
}
