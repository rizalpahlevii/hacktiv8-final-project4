package controller

import "github.com/gin-gonic/gin"

type TransactionControllerContract interface {
	GetTransactions(context *gin.Context)
	GetUserTransactions(context *gin.Context)
	CreateTransaction(context *gin.Context)
}
