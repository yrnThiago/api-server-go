package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
)

type HealthRouter struct {
	Path           string
	HealthHandlers handlers.HealthHandler
}

func NewHealthRouter(healthHandlers *handlers.HealthHandler) *HealthRouter {
	return &HealthRouter{
		Path:           "/health",
		HealthHandlers: *healthHandlers,
	}
}

func (h *HealthRouter) GetRouter() chi.Router {
	chi := chi.NewRouter()
	chi.Get("/ping", h.HealthHandlers.Ping)

	return chi
}
