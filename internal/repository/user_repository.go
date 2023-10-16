package repository

import (
	"github.com/bifidokk/go-api/internal/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

type UserRepository interface {
	FindByEmail(email string) (entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) FindByEmail(email string) (user entity.User, err error) {
	err = ur.database.Where("email = ?", email).First(&user).Error
	
	return user, err
}