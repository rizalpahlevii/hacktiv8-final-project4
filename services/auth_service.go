package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"hacktiv8-final-project4/DTOs"
	"hacktiv8-final-project4/exception"
	"hacktiv8-final-project4/helpers"
	"hacktiv8-final-project4/repositories"
)

type AuthService struct {
	userRepository repositories.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewAuthService(userRepository repositories.UserRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return AuthService{
		userRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service AuthService) Login(input DTOs.LoginDTO) (string, error) {
	err := service.Validate.Struct(input)
	helpers.PanicIfError(err)

	user := service.userRepository.GetUserByEmail(input.Email)

	if user.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("user not found")))
	}

	if !user.VerifyPassword(input.Password) {
		panic(exception.NewLoginError(errors.New("wrong password")))
	}

	return "token", nil
}
