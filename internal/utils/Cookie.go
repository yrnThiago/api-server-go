package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yrnThiago/api-server-go/internal/config"
)

var SECRET_KEY = []byte(config.Env.SECRET_KEY)

func GetCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		fmt.Println("Error getting cookie")
		return "", err
	}
	return cookie.Value, nil
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
