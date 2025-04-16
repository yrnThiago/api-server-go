package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authCookieValue, _ := utils.GetCookie(c, config.Env.COOKIE_NAME)
	userAuthorization, _ := utils.GetFormattedAuthToken(authCookieValue)

	token, err := utils.VerifyJWT(userAuthorization)
	if err != nil {
		return utils.NewErrorInfo("JWT_INVALID_TOKEN", err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims[string(keys.UserIDKey)]

	c.Locals(string(keys.UserIDKey), userID)

	return c.Next()
}
