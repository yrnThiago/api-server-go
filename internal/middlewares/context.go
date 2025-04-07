package middlewares

import (
	"net/http"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"go.uber.org/zap"
)

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
