package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

var Env EnvVariables

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	Env = EnvVariables{
		PORT:        os.Getenv("PORT"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
}

func GetDatabaseUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Env.DB_USERNAME,
		Env.DB_PASSWORD,
		Env.DB_HOST,
		Env.DB_PORT,
		Env.DB_NAME,
	)
}
