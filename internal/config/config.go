package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var EnvDetails map[string]string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	EnvDetails = map[string]string{
		"PORT": os.Getenv("PORT"),
	}
}

func GetEnv(key string) string {
	return EnvDetails[key]
}
