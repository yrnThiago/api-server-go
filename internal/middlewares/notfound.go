package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/config"
)

func NotFoundMiddleware(c *fiber.Ctx) error {
	config.Logger.Warn("invalid endpoint")

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "invalid endpoint"})
}
