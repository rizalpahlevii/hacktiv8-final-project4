package routes

import (
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"hacktiv8-final-project4/config"
	"hacktiv8-final-project4/controllers"
	"hacktiv8-final-project4/repositories"
	"hacktiv8-final-project4/services"
)

func UserRoutes(r *gin.Engine) {
	database := config.DatabaseConnection()
	validator := validator2.New()
	
	userRepository := repositories.NewUserRepository(database)
	userService := services.NewAuthService(userRepository, database, validator)
	authHandler := controllers.NewAuthController(userService)
	router := r.Group("/users")
	{
		router.POST("/login", authHandler.Login)
		router.POST("/register", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Register",
			})
		})
		router.POST("/topup", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Topup",
			})
		})

	}
}
