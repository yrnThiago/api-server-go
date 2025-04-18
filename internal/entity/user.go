package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string
	Password string
	gorm.Model
}

func NewUser(email, password string) *User {
	return &User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
	}
}
