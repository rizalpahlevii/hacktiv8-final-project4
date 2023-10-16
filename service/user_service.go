package service

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"hacktiv8-final-project4/dto"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/request"
)

type UserService struct {
	userRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return UserService{
		userRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service UserService) TopUp(input request.TopUpRequest, loggedUser helper.LoggedUser) dto.TopUpDTO {
	err := service.Validate.Struct(input)
	helper.PanicIfError(err)

	user := service.userRepository.GetUserById(loggedUser.ID)
	service.userRepository.IncreaseBalance(user, input.Balance)

	return dto.NewTopUpDTO(user.Balance)
}
