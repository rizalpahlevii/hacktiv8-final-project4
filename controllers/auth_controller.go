package controllers

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/DTOs"
	"hacktiv8-final-project4/services"
	"net/http"
)

func (handler *AuthController) Register(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

type AuthController struct {
	authService services.AuthService
}

func (handler *AuthController) Login(context *gin.Context) {
	var loginRequest DTOs.LoginDTO
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	token, err := handler.authService.Login(loginRequest)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Login success",
		"token":   token,
	})
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{authService: authService}
}
