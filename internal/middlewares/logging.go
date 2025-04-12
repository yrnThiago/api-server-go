package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
)

func getUserID(c *fiber.Ctx) string {
	userID, ok := c.Locals(string(keys.UserIDKey)).(string)
	if !ok {
		return ""
	}

	return userID
}

func LoggingMiddleware(c *fiber.Ctx) error {
	config.Logger.Info(
		"request received",
		zap.String("request id", c.Locals("requestid").(string)),
		// zap.String("user id", getUserID(c)),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)

	err := c.Next()

	config.Logger.Info(
		"request completed",
		zap.String("request id", c.Locals("requestid").(string)),
		// zap.String("user id", getUserID(c)),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)

	return err
}
