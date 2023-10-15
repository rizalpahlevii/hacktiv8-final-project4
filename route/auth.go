package route

import (
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"hacktiv8-final-project4/config"
	"hacktiv8-final-project4/controller"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/service"
)

func AuthRoutes(r *gin.Engine) {
	database := config.DatabaseConnection()
	validator := validator2.New()

	userRepository := repository.NewUserRepository(database)
	userService := service.NewAuthService(userRepository, database, validator)
	authHandler := controller.NewAuthController(userService)

	router := r.Group("/users")
	{
		router.POST("/login", authHandler.Login)
		router.POST("/register", authHandler.Register)
	}
}
