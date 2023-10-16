package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/request"
	"hacktiv8-final-project4/service"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func (controller UserController) TopUp(context *gin.Context) {
	var topUpRequest request.TopUpRequest
	helper.ReadFromRequestBody(context, &topUpRequest)

	loggedUser := helper.GetLoggedUser(context)
	topUpDTO := controller.userService.TopUp(topUpRequest, loggedUser)
	helper.JsonResponse(context, http.StatusOK, topUpDTO)
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService: userService}
}
