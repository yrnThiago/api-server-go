package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"github.com/yrnThiago/api-server-go/internal/config"
)

var SECRET_KEY = []byte(config.Env.SECRET_KEY)

func GetCookie(c *fiber.Ctx, cookieName string) (string, error) {
	cookie := c.Cookies(cookieName)
	if cookie == "" {
		fmt.Println("Error getting cookie")
		return "", fmt.Errorf("cookie %s not found", cookieName)
	}
	return cookie, nil
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
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
