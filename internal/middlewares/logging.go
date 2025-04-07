package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/config"
	"go.uber.org/zap"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	config.Logger.Info(
		"request received",
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)

	// Processa o próximo handler
	err := c.Next()

	// Após a requisição ser tratada
	config.Logger.Info(
		"request completed",
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Int("status", c.Response().StatusCode()),
	)

	return err
}
