package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"hacktiv8-final-project4/exception"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/request/auth"
)

type AuthService struct {
	userRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return AuthService{
		userRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service AuthService) Login(input auth.LoginRequest) string {
	err := service.Validate.Struct(input)
	helper.PanicIfError(err)

	user := service.userRepository.GetUserByEmail(input.Email)
	fmt.Println(user)
	if user.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("user not found")))
	}

	if !user.VerifyPassword(input.Password) {
		panic(exception.NewLoginError(errors.New("wrong password")))
	}

	return user.GenerateJwtToken()
}
