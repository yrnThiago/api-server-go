package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func ErrorMiddleware(c *fiber.Ctx, err error) error {
	errorInfo := utils.AsErrorInfo(err)

	config.Logger.Warn(errorInfo.GetLowerName(),
		zap.String("request id", c.Locals("requestid").(string)),
		zap.String("error", errorInfo.Message),
	)

	switch errorInfo.Name {
	case "VALIDATION_ERROR":
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorInfo.Message})
	case "RECORD_NOT_FOUND":
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorInfo.Message})
	case "JWT_INVALID_TOKEN":
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": utils.ErrUnauthorizedMsg})
	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": utils.ErrInternalServerMsg})
	}
}
