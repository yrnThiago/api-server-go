package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/entity"
	infra "github.com/yrnThiago/api-server-go/internal/infra/redis"
)

func RateLimitMiddleware(c *fiber.Ctx) error {
	if infra.Redis.IsUp && !infra.Redis.Allow(c.IP()) {
		return entity.ErrRateLimit
	}

	return c.Next()
}
