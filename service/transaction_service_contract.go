package service

import (
	"hacktiv8-final-project4/httpresponse"
	"hacktiv8-final-project4/request"
)

type TransactionServiceContract interface {
	Create(request request.TransactionRequest) httpresponse.CreateTransactionResponse
	GetAll() []httpresponse.AllTransactionResponse
	GetByUser(userId int) []httpresponse.UserTransactionResponse
}
