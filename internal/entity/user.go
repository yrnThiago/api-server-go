package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user *User) error
	GetMany() ([]*User, error)
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	UpdateById(user, newUserBody *User) (*User, error)
	DeleteById(id string) error
}

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
