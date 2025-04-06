package configroutes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/routes"
)

func ProtectedRouter() chi.Router {
	protectedHandlers := handlers.NewProtectedHandlers()
	protectedRouter := routes.NewProtectedRouter(protectedHandlers)

	return protectedRouter.GetRouter()
}
