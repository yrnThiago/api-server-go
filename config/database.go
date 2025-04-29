package config

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type GoEnv string

const (
	LOCAL_ENV        GoEnv = "local"
	DEV_ENV          GoEnv = "dev"
	PROD_ENV         GoEnv = "production"
	tempDatabasePath       = "temp.db"
)

func getDatabaseUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Env.DB_USERNAME,
		Env.DB_PASSWORD,
		Env.DB_HOST,
		Env.DB_PORT,
		Env.DB_NAME,
	)
}

func connectDatabase() {
	var err error
	if Env.GO_ENV == LOCAL_ENV {
		DB, err = gorm.Open(mysql.Open(getDatabaseUrl()), &gorm.Config{})
		if err != nil {
			Logger.Panic("failed to connect to mysql database")
		}

		return
	}

	DB, err = gorm.Open(sqlite.Open(tempDatabasePath), &gorm.Config{})
	if err != nil {
		Logger.Panic("failed to connect to sqlite database")
	}
}

func DatabaseInit() {
	connectDatabase()
	DB.Migrator().AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Order{}, &entity.OrderItems{})

	Logger.Info(
		"db successfully initialized",
		zap.String("env", string(Env.GO_ENV)),
	)
}
