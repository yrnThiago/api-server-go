package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	config.Logger.Info(
		"request received",
		zap.String("request id", c.Locals("requestid").(string)),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)

	// Processa o próximo handler
	err := c.Next()

	// Após a requisição ser tratada
	config.Logger.Info(
		"request completed",
		zap.String("request id", c.Locals("requestid").(string)),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)

	return err
}
