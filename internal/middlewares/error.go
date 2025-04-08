package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	c.Next()

	contextErrorVal := c.Locals(string(keys.ErrorKey))
	if contextErrorVal != nil {
		contextError := contextErrorVal.(*utils.ErrorInfo)

		config.Logger.Info("error occurred",
			zap.Int("status", contextError.StatusCode),
			zap.String("message", contextError.Message),
		)

		return c.Status(contextError.StatusCode).JSON(contextError)
	}

	return nil
}
