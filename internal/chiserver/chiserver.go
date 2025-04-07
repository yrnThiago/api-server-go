package chiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	configroutes "github.com/yrnThiago/api-server-go/internal/config/routes"
	"github.com/yrnThiago/api-server-go/internal/middlewares"
)

func Init() {
	router := chi.NewRouter()

	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.ErrorMiddleware)

	// public routes
	router.Group(func(router chi.Router) {
		router.Mount("/health", configroutes.HealthRouter())
	})

	router.Group(func(router chi.Router) {
		router.Mount("/auth", configroutes.AuthRouter())
	})

	// private routes
	router.Group(func(router chi.Router) {
		router.Use(middlewares.AuthMiddleware)
		router.Use(middlewares.ContextMiddleware)
		router.Mount("/orders", configroutes.OrderRouter())
		router.Mount("/products", configroutes.ProductRouter())
		router.Mount("/protected", configroutes.ProtectedRouter())
	})

	config.Logger.Info(
		"server listening",
		zap.String("port", config.Env.PORT),
	)

	err := http.ListenAndServe(":"+config.Env.PORT, router)
	if err != nil {
		config.Logger.Fatal("server connect failed")
	}
}
