package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
	"go.uber.org/zap"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)

			ctx := r.Context()
			contextError := ctx.Value(keys.ErrorKey)
			if contextError != nil {
				contextError := contextError.(*utils.ErrorInfo)

				switch contextError.StatusCode {
				case http.StatusForbidden:
					w.WriteHeader(http.StatusForbidden)
					json.NewEncoder(w).Encode(contextError)
				default:
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(contextError)
				}

				config.Logger.Info("error occured",
					zap.Int("status", contextError.StatusCode),
					zap.String("message", contextError.Message),
				)
			}
		},
	)
}
