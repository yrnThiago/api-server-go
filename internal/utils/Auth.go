package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yrnThiago/api-server-go/config"
	"golang.org/x/crypto/bcrypt"
)

const UserIdKeyCtx string = "userID"

var secretKey = []byte(config.Env.SecretKey)

func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			UserIdKeyCtx: userID,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}

func ConvertStructToString(obj any) (string, error) {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return string(jsonStr), err
}

func GenerateHash(val string) (string, error) {
	hash := sha256.New()

	// Write the input string as bytes to the hash
	hash.Write([]byte(val))

	// Get the final hashed value as a byte slice
	hashedBytes := hash.Sum(nil)

	// Convert the hashed bytes to a hexadecimal string
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString, nil
}
