package controller

import "github.com/gin-gonic/gin"

type CategoryControllerInterface interface {
	GetCategories(context *gin.Context)
	CreateCategory(context *gin.Context)
	UpdateCategory(context *gin.Context)
	DeleteCategory(context *gin.Context)
}
