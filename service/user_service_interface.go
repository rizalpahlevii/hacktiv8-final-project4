package service

import (
	dto "hacktiv8-final-project4/dto/user"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/request"
)

type UserServiceInterface interface {
	TopUp(request request.TopUpRequest, loggedUser helper.LoggedUser) dto.TopUpDTO
}
