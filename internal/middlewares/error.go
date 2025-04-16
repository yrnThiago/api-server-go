package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	c.Next()

	contextErrorVal := c.Locals(string(keys.ErrorKey))
	if contextErrorVal != nil {
		contextError := contextErrorVal.(*utils.ErrorInfo)

		config.Logger.Info(contextError.Name,
			zap.Int("status", contextError.StatusCode),
			zap.String("message", contextError.Error),
		)

		switch contextError.Name {
		case "ValidationError":
			return c.Status(contextError.StatusCode).JSON(fiber.Map{"error": contextError.Error})
		case fiber.ErrBadRequest.Message:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fiber.ErrBadRequest.Message})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fiber.ErrInternalServerError.Message})
		}
	}

	return nil
}
