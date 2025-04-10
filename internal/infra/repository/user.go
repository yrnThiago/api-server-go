package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/yrnThiago/api-server-go/internal/models"
)

type UserRepositoryMysql struct {
	DB *gorm.DB
}

func NewUserRepositoryMysql(db *gorm.DB) *UserRepositoryMysql {
	return &UserRepositoryMysql{
		DB: db,
	}
}

func (r *UserRepositoryMysql) Add(user *models.User) error {
	res := r.DB.Create(user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *UserRepositoryMysql) GetMany() ([]*models.User, error) {
	var users []*models.User
	res := r.DB.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (r *UserRepositoryMysql) GetById(userID string) (*models.User, error) {
	var user *models.User
	res := r.DB.Limit(1).First(&user, "id = ?", userID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepositoryMysql) GetByEmail(userEmail string) (*models.User, error) {
	var user *models.User
	res := r.DB.Limit(1).First(&user, "email = ?", userEmail)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepositoryMysql) UpdateById(
	userID string,
	newUser *models.User,
) (*models.User, error) {
	user, err := r.GetById(userID)
	if err != nil {
		return nil, err
	}

	r.DB.Model(&user).Omit("ID").Updates(newUser)

	return user, nil
}

func (r *UserRepositoryMysql) DeleteById(userID string) error {
	var user *models.User
	res := r.DB.Delete(&user, "id = ?", userID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
