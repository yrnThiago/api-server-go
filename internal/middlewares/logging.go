package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	config.Logger.Info(
		"request received",
		zap.String("request id", c.Locals("requestid").(string)),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)

	err := c.Next()

	config.Logger.Info(
		"request sent",
		zap.String("request id", c.Locals("requestid").(string)),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)

	return err
}
