package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type AuthRouter struct {
	Path         string
	AuthHandlers *handlers.AuthHandler
}

func NewAuthRouter(authHandlers *handlers.AuthHandler) *AuthRouter {
	return &AuthRouter{
		Path:         "/auth",
		AuthHandlers: authHandlers,
	}
}

func (h *AuthRouter) GetRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/login", h.AuthHandlers.Login)

	return router
}
