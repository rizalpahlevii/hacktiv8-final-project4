package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/request"
	"hacktiv8-final-project4/service"
	"net/http"
)

type TransactionController struct {
	transactionService *service.TransactionService
}

func (controller TransactionController) GetTransactions(context *gin.Context) {
	transactions := controller.transactionService.GetAll()
	helper.JsonResponse(context, http.StatusOK, transactions)
}

func (controller TransactionController) GetUserTransactions(context *gin.Context) {
	transactions := controller.transactionService.GetByUser(1)
	helper.JsonResponse(context, http.StatusOK, transactions)
}

func (controller TransactionController) CreateTransaction(context *gin.Context) {
	var transactionRequest request.TransactionRequest
	helper.ReadFromRequestBody(context, &transactionRequest)
	loggedUser := helper.GetLoggedUser(context)
	transactionRequest.UserId = loggedUser.ID

	transaction := controller.transactionService.Create(transactionRequest)
	helper.JsonResponse(context, http.StatusCreated, transaction)
}

func NewTransactionController(transactionService *service.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}
