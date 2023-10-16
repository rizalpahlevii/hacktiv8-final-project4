package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	TopUp(context *gin.Context)
}
