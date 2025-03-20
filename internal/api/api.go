package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/gdlp-go/internal/config"
	"github.com/yrnThiago/gdlp-go/internal/handlers"
)

type Server struct {
	Logger *slog.Logger
}

type Response struct {
	Message string `json:"message"`
}

var Logger *slog.Logger

func CreateLogger() {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)

	Logger = myslog
}

func CreateServer(
	productHandlers *handlers.ProductHandlers,
	orderHandlers *handlers.OrderHandlers,
) {
	chi := chi.NewRouter()
	// mux := http.NewServeMux()
	Logger.Info("Server listening", "port", config.GetEnv("PORT"))

	setupHandlers(chi, productHandlers, orderHandlers)
	err := http.ListenAndServe(":"+config.GetEnv("PORT"), chi)
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

func ping(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(Response{"pong"})
	if err != nil {
		Logger.Error("Parsing JSON")
		http.Error(w, "Parsing JSON", http.StatusInternalServerError)
	}
}

func setupHandlers(
	chi *chi.Mux,
	productHandlers *handlers.ProductHandlers,
	orderHandlers *handlers.OrderHandlers,
) {
	chi.Use(loggingMiddleware, errorMiddleware)
	chi.Get("/ping", ping)

	chi.Post("/checkout", orderHandlers.Add)
	chi.Get("/orders", orderHandlers.GetMany)
	chi.Get("/orders/{id}", orderHandlers.GetById)
	chi.Post("/orders/{id}", orderHandlers.UpdateById)
	chi.Delete("/orders/{id}", orderHandlers.DeleteById)

	chi.Post("/product", productHandlers.Add)
	chi.Get("/product", productHandlers.GetMany)
	chi.Get("/product/{id}", productHandlers.GetById)
	chi.Post("/product/{id}", productHandlers.UpdateById)
	chi.Delete("/product/{id}", productHandlers.DeleteById)
}
