package middlewares

import (
	"context"
	"net/http"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
	"go.uber.org/zap"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userAuthorization, _ := utils.GetCookie(r, config.Env.COOKIE_NAME)
			err := utils.VerifyJWT(userAuthorization)
			if err != nil {
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
