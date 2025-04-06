package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
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
			ctx := context.WithValue(r.Context(), keys.UserIDKey, "12346")
			next.ServeHTTP(w, r.WithContext(ctx))
			fmt.Println("Private route...")
		},
	)
}
