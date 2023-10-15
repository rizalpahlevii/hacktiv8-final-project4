package controller

import "github.com/gin-gonic/gin"

type ProductHandlerInterface interface {
	GetProducts(context *gin.Context)
	GetProduct(context *gin.Context)
	CreateProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
}

type CategoryHandlerInterface interface {
	GetCategories(context *gin.Context)
	GetCategory(context *gin.Context)
	CreateCategory(context *gin.Context)
	UpdateCategory(context *gin.Context)
	DeleteCategory(context *gin.Context)
}
