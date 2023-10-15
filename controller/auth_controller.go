package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/request/auth"
	"hacktiv8-final-project4/service"
	"net/http"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{authService: authService}
}

func (handler *AuthController) Register(context *gin.Context) {
	var registerRequest auth.RegisterRequest
	helper.ReadFromRequestBody(context, &registerRequest)

	registerDTO := handler.authService.Register(registerRequest)
	helper.JsonResponse(context, http.StatusCreated, registerDTO)
}

func (handler *AuthController) Login(context *gin.Context) {
	var loginRequest auth.LoginRequest
	helper.ReadFromRequestBody(context, &loginRequest)

	loginDTO := handler.authService.Login(loginRequest)
	helper.JsonResponse(context, http.StatusOK, loginDTO)
}
