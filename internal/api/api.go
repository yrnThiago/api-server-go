package api

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yrnThiago/gdlp-go/internal/handlers"
)

type Server struct {
	Logger *slog.Logger
}

type Response struct {
	Message string `json:"message"`
}

var (
	Logger *slog.Logger
	PORT   string
)

func CreateLogger() {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)

	Logger = myslog
}

func CreateServer(productHandlers *handlers.ProductHandlers, orderHandlers *handlers.OrderHandlers) {
	mux := http.NewServeMux()
	Logger.Info("Server listening", "port", PORT)

	setupHandlers(mux, productHandlers, orderHandlers)
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		Logger.Error("Servidor deu merda!")
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
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
	m *http.ServeMux,
	productHandlers *handlers.ProductHandlers,
	orderHandlers *handlers.OrderHandlers,
) {
	m.Handle("/ping", loggingMiddleware(errorMiddleware(http.HandlerFunc(ping))))
	m.Handle(
		"/checkout",
		loggingMiddleware(errorMiddleware(http.HandlerFunc(orderHandlers.CreateOrderHandler))),
	)
	m.Handle(
		"/orders",
		loggingMiddleware(errorMiddleware(http.HandlerFunc(orderHandlers.ListOrderHandler))),
	)
	m.Handle(
		"/products",
		loggingMiddleware(errorMiddleware(http.HandlerFunc(productHandlers.ListProductsHandler))),
	)
	m.Handle(
		"/addproduct",
		loggingMiddleware(errorMiddleware(http.HandlerFunc(productHandlers.CreateProductHandler))),
	)
}
