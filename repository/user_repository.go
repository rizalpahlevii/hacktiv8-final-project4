package repository

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (repository *UserRepository) CreateUser(input request.RegisterRequest) model.User {
	user := model.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
	}
	repository.DB.Create(&user)
	return user
}

func (repository *UserRepository) IncreaseBalance(user model.User, amount int) {
	user.Balance += amount
	repository.DB.Save(&user)
}

func (repository *UserRepository) GetUserByEmail(email string) model.User {
	var user model.User
	repository.DB.Where("email = ?", email).First(&user)
	return user
}

func (repository *UserRepository) GetUserById(id uint) model.User {
	var user model.User
	repository.DB.Where("id = ?", id).First(&user)
	return user
}
