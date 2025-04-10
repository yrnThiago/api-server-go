package config

import (
	"fmt"

	"github.com/yrnThiago/api-server-go/internal/models"
	"gorm.io/driver/mysql"
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
	db, err := gorm.Open(mysql.Open(getDatabaseUrl()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItems{})
	DB = db
}
