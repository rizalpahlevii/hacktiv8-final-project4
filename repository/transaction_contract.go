package repository

import (
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type TransactionContract interface {
	Create(request request.TransactionRequest) model.TransactionHistory
	All() []model.TransactionHistory
	GetByUserId(userId int) []model.TransactionHistory
}
