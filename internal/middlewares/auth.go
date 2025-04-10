package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authCookieValue, _ := utils.GetCookie(c, config.Env.COOKIE_NAME)
	userAuthorization, _ := utils.GetFormattedAuthToken(authCookieValue)

	fmt.Println(authCookieValue, userAuthorization)

	err := utils.VerifyJWT(userAuthorization)
	if err != nil {
		errorInfo := utils.NewErrorInfo(http.StatusForbidden, "access denied")
		c.Locals(string(keys.ErrorKey), errorInfo)
		return errors.New(errorInfo.Message)
	}

	c.Locals(string(keys.UserIDKey), userAuthorization)

	config.Logger.Info(
		"access granted",
		zap.String("user id", userAuthorization),
	)

	return c.Next()
}
