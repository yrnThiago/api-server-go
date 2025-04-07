package middlewares

import (
	"net/http"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)

			ctx := r.Context()
			contextError := ctx.Value(keys.ErrorKey)
			if contextError != nil {
				codeError := contextError.(int)

				switch codeError {
				case http.StatusForbidden:
					http.Error(w, "access denied", http.StatusForbidden)
				default:
					http.Error(w, "something went wrong", http.StatusForbidden)
				}

				config.Logger.Info("error middleware catch error")
			}
		},
	)
}
