package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	DB_USERNAME       string
	DB_PASSWORD       string
	DB_HOST           string
	DB_PORT           string
	DB_NAME           string
	PORT              string
	GO_ENV            GoEnv
	SECRET_KEY        string
	SKIP_AUTH         bool
	COOKIE_NAME       string
	COOKIE_EXPIRES_AT time.Duration
	NEW_ORDERS_TOPIC  string
	RDB_ADDRESS       string
	RDB_PASSWORD      string
	RDB_DB            int
	RATE_LIMIT        int
	RATE_LIMIT_WINDOW time.Duration
	LOGS_FILE_NAME    string
}

var Env EnvVariables

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(".env missing")
	}

	skipAuth, _ := strconv.ParseBool(os.Getenv("SKIP_AUTH"))
	rdbDb, _ := strconv.Atoi(os.Getenv("RDB_DB"))
	rateLimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	rateLimitWindow, _ := time.ParseDuration(os.Getenv("RATE_LIMIT_WINDOW"))
	cookieExpiresAt, _ := time.ParseDuration(os.Getenv("COOKIE_EXPIRES_AT"))

	Env = EnvVariables{
		DB_USERNAME:       os.Getenv("DB_USERNAME"),
		DB_PASSWORD:       os.Getenv("DB_PASSWORD"),
		DB_HOST:           os.Getenv("DB_HOST"),
		DB_PORT:           os.Getenv("DB_PORT"),
		DB_NAME:           os.Getenv("DB_NAME"),
		PORT:              os.Getenv("PORT"),
		GO_ENV:            GoEnv(os.Getenv("GO_ENV")),
		SECRET_KEY:        os.Getenv("SECRET_KEY"),
		SKIP_AUTH:         skipAuth,
		COOKIE_NAME:       os.Getenv("COOKIE_NAME"),
		COOKIE_EXPIRES_AT: cookieExpiresAt,
		NEW_ORDERS_TOPIC:  os.Getenv("NEW_ORDERS_TOPIC"),
		RDB_ADDRESS:       os.Getenv("RDB_ADDRESS"),
		RDB_PASSWORD:      os.Getenv("RDB_PASSWORD"),
		RDB_DB:            rdbDb,
		RATE_LIMIT:        rateLimit,
		RATE_LIMIT_WINDOW: rateLimitWindow,
		LOGS_FILE_NAME:    os.Getenv("LOGS_FILE_NAME"),
	}
}
