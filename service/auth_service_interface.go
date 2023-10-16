package service

import (
	dto "hacktiv8-final-project4/dto/auth"
	"hacktiv8-final-project4/request"
)

type AuthServiceInterface interface {
	Login(request request.LoginRequest) dto.LoginDTO
	Register(request request.RegisterRequest) dto.RegisterDTO
}
