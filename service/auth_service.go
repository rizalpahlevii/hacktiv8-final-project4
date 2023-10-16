package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"hacktiv8-final-project4/dto"
	"hacktiv8-final-project4/exception"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/request"
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

func (service AuthService) Login(input request.LoginRequest) dto.LoginDTO {
	err := service.Validate.Struct(input)
	helper.PanicIfError(err)

	user := service.userRepository.GetUserByEmail(input.Email)
	if user.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("user not found")))
	}

	if !user.VerifyPassword(input.Password) {
		panic(exception.NewLoginError(errors.New("wrong password")))
	}

	return dto.NewLoginDTO(user.GenerateJwtToken())
}

func (service AuthService) Register(input request.RegisterRequest) dto.RegisterDTO {
	err := service.Validate.Struct(input)
	helper.PanicIfError(err)

	user := service.userRepository.GetUserByEmail(input.Email)
	if user.ID != 0 {
		panic(exception.NewBadRequestError(errors.New("email already exist")))
	}

	user = service.userRepository.CreateUser(input)
	return dto.NewRegisterDTO(user)
}
