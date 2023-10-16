package httpresponse

import (
	"hacktiv8-final-project4/model"
	"time"
)

type TransactionResponse struct {
	ID         int                        `json:"id"`
	ProductID  int                        `json:"product_id"`
	UserID     int                        `json:"user_id"`
	Quantity   int                        `json:"quantity"`
	TotalPrice int                        `json:"total_price"`
	Product    TransactionProductResponse `json:"product"`
}

func NewTransactionResponse(transaction model.TransactionHistory) TransactionResponse {
	return TransactionResponse{
		ID:         transaction.ID,
		ProductID:  transaction.ProductID,
		UserID:     transaction.UserID,
		Quantity:   transaction.Quantity,
		TotalPrice: transaction.TotalPrice,
		Product:    NewTransactionProductResponse(transaction.Product),
	}
}

type TransactionProductResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAT  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewTransactionProductResponse(product model.Product) TransactionProductResponse {
	return TransactionProductResponse{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAT:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

type CreateTransactionResponse struct {
	Message         string `json:"message"`
	TransactionBill struct {
		TotalPrice   int    `json:"total_price"`
		Quantity     int    `json:"quantity"`
		ProductTitle string `json:"product_title"`
	} `json:"transaction_bill"`
}

func NewCreateTransactionResponse(transaction model.TransactionHistory) CreateTransactionResponse {
	product := transaction.Product
	return CreateTransactionResponse{
		Message: "You have successfully purchase the product",
		TransactionBill: struct {
			TotalPrice   int    `json:"total_price"`
			Quantity     int    `json:"quantity"`
			ProductTitle string `json:"product_title"`
		}{
			TotalPrice:   transaction.TotalPrice,
			Quantity:     transaction.Quantity,
			ProductTitle: product.Title,
		},
	}
}

type AllTransactionResponse struct {
	TransactionResponse
	User UserResponse `json:"user"`
}

func NewAllTransactionResponse(transaction model.TransactionHistory) AllTransactionResponse {
	return AllTransactionResponse{
		TransactionResponse: NewTransactionResponse(transaction),
		User:                NewUserResponse(transaction.User),
	}
}

type UserTransactionResponse struct {
	TransactionResponse
}

func NewUserTransactionResponse(transaction model.TransactionHistory) UserTransactionResponse {
	return UserTransactionResponse{
		TransactionResponse: NewTransactionResponse(transaction),
	}
}
