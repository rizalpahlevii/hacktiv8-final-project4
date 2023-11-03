package route

import (
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"hacktiv8-final-project4/config"
	"hacktiv8-final-project4/controller"
	"hacktiv8-final-project4/middleware"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/service"
)

func UserRoutes(r *gin.Engine) {
	database := config.DatabaseConnection()
	validator := validator2.New()

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository, database, validator)
	userController := controller.NewUserController(userService)

	router := r.Group("/users")
	{
		router.Use(middleware.JwtMiddleware)
		router.PATCH("/topup", userController.TopUp)

	}
}
