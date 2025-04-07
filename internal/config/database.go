package config

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/yrnThiago/api-server-go/internal/models"
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
		panic("failed to connect database")
	}

	db.Migrator().AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderItems{})
	DB = db
}
