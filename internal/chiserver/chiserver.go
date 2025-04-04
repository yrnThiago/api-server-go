package chiserver

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/config"
	configroutes "github.com/yrnThiago/api-server-go/internal/config/routes"
)

type Server struct {
	Logger *slog.Logger
}

var Logger *slog.Logger

func CreateLogger() {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)

	Logger = myslog
}

func CreateServer() {
	chi := chi.NewRouter()
	Logger.Info("Server listening", "port", config.Env.PORT)

	setupHandlers(chi)
	err := http.ListenAndServe(":"+config.Env.PORT, chi)
	if err != nil {
		Logger.Error("Servidor deu merda!")
	}
}

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			fmt.Println("Testando group...")
		},
	)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Info("Request received", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
		Logger.Info("Request completed", "method", r.Method, "path", r.URL.Path)
	})
}

// To do auth with context
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			fmt.Println("Private route...")
		},
	)
}

func setupHandlers(
	router *chi.Mux,
) {

	router.Use(loggingMiddleware)

	// public routes
	router.Group(func(router chi.Router) {
		router.Mount("/health", configroutes.HealthRouter())
	})

	// private routes
	router.Group(func(router chi.Router) {
		router.Use(authMiddleware)
		router.Mount("/orders", configroutes.OrderRouter())
		router.Mount("/products", configroutes.ProductRouter())
	})
}
