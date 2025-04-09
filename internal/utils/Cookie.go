package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"github.com/yrnThiago/api-server-go/config"
	"go.uber.org/zap"
)

var SECRET_KEY = []byte(config.Env.SECRET_KEY)

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

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": "testeusern",
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
