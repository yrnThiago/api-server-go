package middlewares

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
)

var errorMessage string

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)

			ctx := r.Context()
			contextError := ctx.Value(keys.ErrorKey)
			if contextError != nil {
				errorStatusCode := contextError.(int)

				if errorStatusCode == http.StatusForbidden {
					errorMessage = "access denied"
				}
				config.Logger.Info(
					errorMessage,
				)

				http.Error(w, errorMessage, errorStatusCode)
				return
			}
		},
	)
}

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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userAuthorization := r.Header.Get("Authorization")
			if userAuthorization == "" {
				ctx := context.WithValue(r.Context(), keys.ErrorKey, http.StatusForbidden)
				*r = *r.WithContext(ctx)
				return
			}

			ctx := context.WithValue(r.Context(), keys.UserIDKey, userAuthorization)
			next.ServeHTTP(w, r.WithContext(ctx))

			config.Logger.Info(
				"access granted",
				zap.String("user id", ctx.Value(keys.UserIDKey).(string)),
			)
		},
	)
}

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userId, _ := ctx.Value(keys.UserIDKey).(string)

			config.Logger.Info(
				"ctx",
				zap.String("user id: ", userId),
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
