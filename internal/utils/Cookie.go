package utils

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/config"
	"go.uber.org/zap"
)

type Option func(c fiber.Cookie) fiber.Cookie

var SECRET_KEY = []byte(config.Env.SECRET_KEY)
var BEARER_KEY = "Bearer "

func GetCookie(c *fiber.Ctx, cookieName string) (string, error) {
	cookie := c.Cookies(cookieName)
	if cookie == "" {
		return "", fmt.Errorf("cookie %s not found", cookieName)
	}
	return cookie, nil
}

func SetCookie(c *fiber.Ctx, cookie *fiber.Cookie) (string, error) {
	c.Cookie(cookie)
	newCookie, err := GetCookie(c, cookie.Name)

	config.Logger.Info(
		"set cookie",
		zap.String("name", cookie.Name),
		zap.String("value", cookie.Value),
	)
	return newCookie, err
}

func SetBearerCookie(c *fiber.Ctx, cookie *fiber.Cookie) (string, error) {
	c.Cookie(cookie)
	newCookie, err := GetCookie(c, cookie.Name)

	config.Logger.Info(
		"set cookie",
		zap.String("name", cookie.Name),
		zap.String("value", cookie.Value),
	)
	return newCookie, err
}
func GetFormattedAuthToken(token string) (string, error) {
	return strings.Replace(token, BEARER_KEY, "", 1), nil
}
