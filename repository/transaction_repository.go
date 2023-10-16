package repository

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/model"
	"hacktiv8-final-project4/request"
)

type TransactionRepository struct {
	*gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: DB}
}

func (repository *TransactionRepository) Create(request request.TransactionRequest) model.TransactionHistory {
	transaction := model.TransactionHistory{
		ProductID:  request.ProductId,
		Quantity:   request.Quantity,
		UserID:     request.UserId,
		TotalPrice: request.TotalPrice,
	}

	repository.DB.Create(&transaction)
	return transaction
}

func (repository *TransactionRepository) All() []model.TransactionHistory {
	var transactions []model.TransactionHistory
	repository.DB.Preload("Product").Preload("User").Find(&transactions)
	return transactions
}

func (repository *TransactionRepository) GetByUserId(userId int) []model.TransactionHistory {
	var transactions []model.TransactionHistory
	repository.DB.Preload("Product").Preload("User").Where("user_id = ?", userId).Find(&transactions)
	return transactions
}
