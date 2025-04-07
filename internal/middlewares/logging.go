package middlewares

import (
	"net/http"

	"github.com/yrnThiago/api-server-go/internal/config"
	"go.uber.org/zap"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config.Logger.Info(
			"request received",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)

		next.ServeHTTP(w, r)

		config.Logger.Info(
			"request completed",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)
	})
}
