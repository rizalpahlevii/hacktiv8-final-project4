package controller

import "github.com/gin-gonic/gin"

type CategoryHandlerInterface interface {
	GetCategories(context *gin.Context)
	GetCategory(context *gin.Context)
	CreateCategory(context *gin.Context)
	UpdateCategory(context *gin.Context)
	DeleteCategory(context *gin.Context)
}
