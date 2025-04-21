package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func RateLimitMiddleware(c *fiber.Ctx) error {
	if config.Redis.IsUp && !config.Redis.Allow(c.IP()) {
		return utils.ErrRateLimit
	}

	return c.Next()
}
