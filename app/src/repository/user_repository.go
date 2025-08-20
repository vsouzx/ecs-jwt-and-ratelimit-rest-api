package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/model"
)

type UserRepositoryInterface interface {
	Create(user *model.User) error
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil
}
