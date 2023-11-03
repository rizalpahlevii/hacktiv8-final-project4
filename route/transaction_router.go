package route

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/config"
	"hacktiv8-final-project4/controller"
	"hacktiv8-final-project4/middleware"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/service"
)

func TransactionRoutes(r *gin.Engine) {
	database := config.DatabaseConnection()
	transactionRepository := repository.NewTransactionRepository(database)
	userRepository := repository.NewUserRepository(database)
	productRepository := repository.NewProductRepository(database)
	categoryRepository := repository.NewCategoryRepository(database)
	transactionService := service.NewTransactionService(transactionRepository, productRepository, &userRepository, categoryRepository)
	transactionController := controller.NewTransactionController(transactionService)
	router := r.Group("/transactions")
	{
		router.Use(middleware.JwtMiddleware)
		router.GET("/user-transactions", middleware.IsAdminMiddleware, transactionController.GetTransactions)
		router.GET("/my-transactions", transactionController.GetUserTransactions)
		router.POST("/", transactionController.CreateTransaction)

	}
}
