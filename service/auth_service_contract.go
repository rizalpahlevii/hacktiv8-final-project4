package service

import (
	"hacktiv8-final-project4/httpresponse"
	"hacktiv8-final-project4/request"
)

type AuthServiceContract interface {
	Login(request request.LoginRequest) httpresponse.LoginResponse
	Register(request request.RegisterRequest) httpresponse.RegisterResponse
}
