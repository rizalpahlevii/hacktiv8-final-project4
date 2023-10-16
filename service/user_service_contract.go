package service

import (
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/httpresponse"
	"hacktiv8-final-project4/request"
)

type UserServiceContract interface {
	TopUp(request request.TopUpRequest, loggedUser helper.LoggedUser) httpresponse.TopUpResponse
}
