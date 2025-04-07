package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	userAuthorization, _ := utils.GetCookie(c, config.Env.COOKIE_NAME)

	err := utils.VerifyJWT(userAuthorization)
	if err != nil {
		errorInfo := utils.NewErrorInfo(http.StatusForbidden, "access denied")
		// Armazenando erro no contexto para middlewares futuros, se necessário
		c.Locals(string(keys.ErrorKey), errorInfo)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": errorInfo.Message,
		})
	}

	// Armazenando o userAuthorization (user ID) no contexto
	c.Locals(string(keys.UserIDKey), userAuthorization)

	config.Logger.Info(
		"access granted",
		zap.String("user id", userAuthorization),
	)

	return c.Next()
}
