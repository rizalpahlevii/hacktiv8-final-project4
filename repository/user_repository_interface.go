package repository

import (
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type UserRepositoryInterface interface {
	CreateUser(input request.RegisterRequest) model.User
	IncreaseBalance(user model.User, amount int) model.User
	DecreaseBalance(user model.User, amount int)
	FindByEmail(email string) model.User
	FindById(id int) model.User
}
