package utils

import (
	"fmt"
	"strings"
	"time"

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

func SetBearerCookie(c *fiber.Ctx, token string) (string, error) {

	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = BEARER_KEY + token
	cookie.Expires = time.Now().Add(config.Env.COOKIE_EXPIRES_AT)
	cookie.Secure = false
	cookie.HTTPOnly = true
	cookie.Path = "/"

	c.Cookie(cookie)
	newCookie, err := GetCookie(c, cookie.Name)

	config.Logger.Info(
		"set cookie",
		zap.String("name", cookie.Name),
		zap.String("value", cookie.Value),
	)

	return newCookie, err
}

func ClearBearerCookie(c *fiber.Ctx) error {
	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Expires = time.Now().Add(-3 * time.Second)
	cookie.HTTPOnly = true

	_, err := SetCookie(c, cookie)
	return err

}

func GetFormattedAuthToken(token string) (string, error) {
	return strings.Replace(token, BEARER_KEY, "", 1), nil
}
