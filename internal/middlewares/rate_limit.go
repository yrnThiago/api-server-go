package middlewares

import (
	"github.com/gofiber/fiber/v2"
	infra "github.com/yrnThiago/api-server-go/internal/infra/redis"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func RateLimitMiddleware(c *fiber.Ctx) error {
	if infra.Redis.IsUp && !infra.Redis.Allow(c.IP()) {
		return utils.ErrRateLimit
	}

	return c.Next()
}
