package controller

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}
