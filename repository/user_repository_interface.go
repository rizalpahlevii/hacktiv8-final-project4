package repository

import (
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type UserRepositoryInterface interface {
	CreateUser(input request.RegisterRequest) model.User
	IncreaseBalance(user model.User, amount int)
	GetUserByEmail(email string) model.User
	GetUserById(id uint) model.User
}
