package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
)

func ContextMiddleware(c *fiber.Ctx) error {
	userId, _ := c.Locals(keys.UserIDKey).(string)

	config.Logger.Info(
		"ctx",
		zap.String("user id: ", userId),
	)

	return c.Next()
}
