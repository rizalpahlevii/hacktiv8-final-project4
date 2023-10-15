package repositories

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/models"
)

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (repository *UserRepository) GetUserByEmail(email string) models.User {
	var user models.User
	repository.DB.Where("email = ?", email).First(&user)
	return user
}
