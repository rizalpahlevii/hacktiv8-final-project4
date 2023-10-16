package service

import (
	"errors"
	"hacktiv8-final-project4/exception"
	"hacktiv8-final-project4/httpresponse"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/request"
)

type TransactionService struct {
	transactionRepository *repository.TransactionRepository
	productRepository     *repository.ProductRepository
	userRepository        *repository.UserRepository
	categoryRepository    *repository.CategoryRepository
}

func (service TransactionService) Create(request request.TransactionRequest) httpresponse.CreateTransactionResponse {
	product := service.productRepository.FindById(request.ProductId)

	if product.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("product not found")))
	}

	if product.Stock < request.Quantity {
		panic(exception.NewBadRequestError(errors.New("product stock is not enough")))
	}

	request.TotalPrice = product.Price * request.Quantity

	user := service.userRepository.FindById(request.UserId)
	if user.ID == 0 {
		panic(exception.NewNotFoundError(errors.New("user not found")))
	}

	if user.Balance < request.TotalPrice {
		panic(exception.NewBadRequestError(errors.New("balance is not enough")))
	}

	transaction := service.transactionRepository.Create(request)
	if transaction.ID == 0 {
		panic(errors.New("failed to create transaction"))
	}

	service.productRepository.DecreaseStock(product, request.Quantity)
	service.userRepository.DecreaseBalance(user, request.TotalPrice)
	service.categoryRepository.IncreaseSoldProductAmount(product.Category, request.Quantity)
	service.transactionRepository.DB.Preload("Product").First(&transaction, transaction.ID)

	return httpresponse.NewCreateTransactionResponse(transaction)
}

func (service TransactionService) GetAll() []httpresponse.AllTransactionResponse {
	transactions := service.transactionRepository.All()
	var transactionResponses []httpresponse.AllTransactionResponse
	for _, value := range transactions {
		transactionResponses = append(transactionResponses, httpresponse.NewAllTransactionResponse(value))
	}

	if len(transactionResponses) == 0 {
		return []httpresponse.AllTransactionResponse{}
	}

	return transactionResponses
}

func (service TransactionService) GetByUser(userId int) []httpresponse.UserTransactionResponse {
	transactions := service.transactionRepository.GetByUserId(userId)
	var transactionResponses []httpresponse.UserTransactionResponse
	for _, value := range transactions {
		transactionResponses = append(transactionResponses, httpresponse.NewUserTransactionResponse(value))
	}

	if len(transactionResponses) == 0 {
		return []httpresponse.UserTransactionResponse{}
	}

	return transactionResponses
}

func NewTransactionService(transactionRepository *repository.TransactionRepository, productRepository *repository.ProductRepository, userRepository *repository.UserRepository, categoryRepository *repository.CategoryRepository) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		productRepository:     productRepository,
		userRepository:        userRepository,
		categoryRepository:    categoryRepository,
	}
}
