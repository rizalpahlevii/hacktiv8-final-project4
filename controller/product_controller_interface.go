package controller

import "github.com/gin-gonic/gin"

type ProductControllerInterface interface {
	GetProducts(context *gin.Context)
	GetProduct(context *gin.Context)
	CreateProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
}
