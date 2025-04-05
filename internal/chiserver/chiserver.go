package chiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/config"
	configroutes "github.com/yrnThiago/api-server-go/internal/config/routes"
	"github.com/yrnThiago/api-server-go/internal/middlewares"
	"go.uber.org/zap"
)

func Init() {
	router := chi.NewRouter()

	router.Use(middlewares.LoggingMiddleware)

	// public routes
	router.Group(func(router chi.Router) {
		router.Mount("/health", configroutes.HealthRouter())
	})

	// private routes
	router.Group(func(router chi.Router) {
		router.Use(middlewares.AuthMiddleware)
		router.Mount("/orders", configroutes.OrderRouter())
		router.Mount("/products", configroutes.ProductRouter())
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
