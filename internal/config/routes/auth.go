package configroutes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/routes"
)

func AuthRouter() chi.Router {
	authHandlers := handlers.NewAuthHandlers()
	authRouter := routes.NewAuthRouter(authHandlers)

	return authRouter.GetRouter()
}
