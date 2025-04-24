package config

import (
	"fmt"

	"github.com/yrnThiago/api-server-go/internal/entity"
	"go.uber.org/zap"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDatabaseUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Env.DB_USERNAME,
		Env.DB_PASSWORD,
		Env.DB_HOST,
		Env.DB_PORT,
		Env.DB_NAME,
	)
}

func DatabaseInit() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		Logger.Panic("failed to connect to database")
	}

	db.Migrator().AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Order{}, &entity.OrderItems{})
	DB = db

	Logger.Info(
		"db successfully initialized",
		zap.String("host", Env.DB_HOST),
		zap.String("port", Env.DB_PORT),
		zap.String("db name", Env.DB_NAME),
	)
}
