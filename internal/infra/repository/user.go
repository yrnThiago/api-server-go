package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type UserRepositoryMysql struct {
	DB *gorm.DB
}

func NewUserRepositoryMysql(db *gorm.DB) *UserRepositoryMysql {
	return &UserRepositoryMysql{
		DB: db,
	}
}

func (r *UserRepositoryMysql) Add(user *entity.User) error {
	res := r.DB.Create(user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *UserRepositoryMysql) GetMany() ([]*entity.User, error) {
	var users []*entity.User
	res := r.DB.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (r *UserRepositoryMysql) GetById(userID string) (*entity.User, error) {
	var user *entity.User
	res := r.DB.Limit(1).First(&user, "id = ?", userID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrUserNotFound
		}

		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepositoryMysql) GetByLogin(userEmail string) (*entity.User, error) {
	var user *entity.User
	res := r.DB.Limit(1).First(&user, "email = ?", userEmail)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepositoryMysql) UpdateById(user *entity.User) (*entity.User, error) {
	res := r.DB.Save(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepositoryMysql) DeleteById(userID string) error {
	var user *entity.User
	res := r.DB.Delete(&user, "id = ?", userID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
