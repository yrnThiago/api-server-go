package middlewares

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"

	"github.com/yrnThiago/api-server-go/internal/config"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			fmt.Println("Testando group...")
		},
	)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config.Logger.Info(
			"Request received",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)

		next.ServeHTTP(w, r)

		config.Logger.Info(
			"Request completed",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)
	})
}

// To do auth with context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			fmt.Println("Private route...")
		},
	)
}
