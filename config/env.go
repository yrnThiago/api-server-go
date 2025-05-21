package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	DBUsername           string
	DBPassword           string
	DBHost               string
	DBPort               string
	DBName               string
	Port                 string
	GoEnv                GoEnv
	SecretKey            string
	SkipAuth             bool
	CookieName           string
	CookieExpiresAt      time.Duration
	NatsURL              string
	RdbAddress           string
	RdbPassword          string
	RdbDB                int
	RateLimit            int
	RateLimitWindow      time.Duration
	UserSessionExpiresAt time.Duration
	OfferExpiresAt       time.Duration
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
	userSessionExpiresAt, _ := time.ParseDuration(os.Getenv("USER_SESSION_EXPIRES_AT"))
	offerExpiresAt, _ := time.ParseDuration(os.Getenv("OFFER_EXPIRES_AT"))

	Env = EnvVariables{
		DBUsername:           os.Getenv("DB_USERNAME"),
		DBPassword:           os.Getenv("DB_PASSWORD"),
		DBHost:               os.Getenv("DB_HOST"),
		DBPort:               os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		Port:                 os.Getenv("PORT"),
		GoEnv:                GoEnv(os.Getenv("GO_ENV")),
		SecretKey:            os.Getenv("SECRET_KEY"),
		SkipAuth:             skipAuth,
		CookieName:           os.Getenv("COOKIE_NAME"),
		CookieExpiresAt:      cookieExpiresAt,
		NatsURL:              os.Getenv("NATS_URL"),
		RdbAddress:           os.Getenv("RDB_ADDRESS"),
		RdbPassword:          os.Getenv("RDB_PASSWORD"),
		RdbDB:                rdbDb,
		RateLimit:            rateLimit,
		RateLimitWindow:      rateLimitWindow,
		UserSessionExpiresAt: userSessionExpiresAt,
		OfferExpiresAt:       offerExpiresAt,
	}
}
