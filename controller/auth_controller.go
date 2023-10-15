package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/dto"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/request/auth"
	"hacktiv8-final-project4/service"
	"net/http"
)

func (handler *AuthController) Register(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

type AuthController struct {
	authService service.AuthService
}

func (handler *AuthController) Login(context *gin.Context) {
	var loginRequest auth.LoginRequest
	helper.ReadFromRequestBody(context, &loginRequest)

	token := handler.authService.Login(loginRequest)
	helper.JsonResponse(context, http.StatusOK, dto.LoginDTO{Token: token})
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{authService: authService}
}
