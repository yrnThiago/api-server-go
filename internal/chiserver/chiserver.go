package chiserver

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
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

func CreateServer(
	healthHandlers *handlers.HealthHandler,
	productHandlers *handlers.ProductHandlers,
	orderHandlers *handlers.OrderHandlers,
) {
	chi := chi.NewRouter()
	Logger.Info("Server listening", "port", config.Env.PORT)

	setupHandlers(chi, healthHandlers, productHandlers, orderHandlers)
	err := http.ListenAndServe(":"+config.Env.PORT, chi)
	if err != nil {
		Logger.Error("Servidor deu merda!")
	}
}

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
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

func setupHandlers(
	chi *chi.Mux,
	healthHandlers *handlers.HealthHandler,
	productHandlers *handlers.ProductHandlers,
	orderHandlers *handlers.OrderHandlers,
) {
	chi.Use(loggingMiddleware, errorMiddleware)
	chi.Get("/ping", healthHandlers.Ping)

	chi.Post("/checkout", orderHandlers.Add)
	chi.Get("/orders", orderHandlers.GetMany)
	chi.Get("/orders/{id}", orderHandlers.GetById)
	chi.Put("/orders/{id}", orderHandlers.UpdateById)
	chi.Delete("/orders/{id}", orderHandlers.DeleteById)

	chi.Post("/products", productHandlers.Add)
	chi.Get("/products", productHandlers.GetMany)
	chi.Get("/products/{id}", productHandlers.GetById)
	chi.Put("/products/{id}", productHandlers.UpdateById)
	chi.Delete("/products/{id}", productHandlers.DeleteById)
}
