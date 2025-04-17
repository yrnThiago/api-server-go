package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func ErrorMiddleware(c *fiber.Ctx, err error) error {
	customError, ok := err.(*utils.ErrorInfo)
	if !ok {
		customError = utils.GetInternalError()
	}

	config.Logger.Warn(customError.Name,
		zap.String("request id", c.Locals("requestid").(string)),
		zap.String("error", customError.Message),
	)

	switch customError.Name {
	case "ValidationError":
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": customError.Message})
	case "JWT_INVALID_TOKEN":
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": customError.Message})
	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": fiber.ErrInternalServerError.Message})
	}
}
