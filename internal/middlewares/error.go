package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/entity"
)

func ErrorMiddleware(c *fiber.Ctx, err error) error {
	errorInfo := entity.AsErrorInfo(err)

	config.Logger.Warn(errorInfo.GetLowerName(),
		zap.String("request id", c.Locals("requestid").(string)),
		zap.String("error", errorInfo.Message),
	)

	switch errorInfo.Name {
	case entity.ErrValidationName:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorInfo.Errors})
	case entity.ErrRecordNotFoundName:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorInfo.Message})
	case entity.ErrInvalidJwtTokenName:
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": errorInfo.Message})
	case entity.ErrRateLimitName:
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": errorInfo.Message})
	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": entity.ErrInternalServerMsg})
	}
}
