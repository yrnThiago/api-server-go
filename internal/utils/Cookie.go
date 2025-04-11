package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

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
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userID,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return SECRET_KEY, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func GetFormattedAuthToken(token string) (string, error) {
	return strings.Replace(token, BEARER_KEY, "", 1), nil
}
