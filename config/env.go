package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	PORT             string
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_HOST          string
	DB_PORT          string
	DB_NAME          string
	COOKIE_NAME      string
	SECRET_KEY       string
	NEW_ORDERS_TOPIC string
	RDB_ADDRESS      string
	RDB_PASSWORD     string
	RDB_DB           string
	LOGS_FILE_NAME   string
}

var Env EnvVariables

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(".env missing")
	}

	Env = EnvVariables{
		PORT:             os.Getenv("PORT"),
		DB_USERNAME:      os.Getenv("DB_USERNAME"),
		DB_PASSWORD:      os.Getenv("DB_PASSWORD"),
		DB_HOST:          os.Getenv("DB_HOST"),
		DB_PORT:          os.Getenv("DB_PORT"),
		DB_NAME:          os.Getenv("DB_NAME"),
		COOKIE_NAME:      os.Getenv("COOKIE_NAME"),
		SECRET_KEY:       os.Getenv("SECRET_KEY"),
		NEW_ORDERS_TOPIC: os.Getenv("NEW_ORDERS_TOPIC"),
		RDB_ADDRESS:      os.Getenv("RDB_ADDRESS"),
		RDB_PASSWORD:     os.Getenv("RDB_PASSWORD"),
		RDB_DB:           os.Getenv("RDB_DB"),
		LOGS_FILE_NAME:   os.Getenv("LOGS_FILE_NAME"),
	}
}
