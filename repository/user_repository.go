package repository

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/model"
)

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (repository *UserRepository) GetUserByEmail(email string) model.User {
	var user model.User
	repository.DB.Where("email = ?", email).First(&user)
	return user
}
