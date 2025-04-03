package configroutes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/routes"
)

func HealthRouter() chi.Router {
	healthHandlers := handlers.NewHealthHandlers()
	healthRouter := routes.NewHealthRouter(healthHandlers)

	return healthRouter.GetRouter()
}
