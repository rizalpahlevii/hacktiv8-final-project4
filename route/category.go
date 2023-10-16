package route

import (
	"github.com/gin-gonic/gin"
	validatorV10 "github.com/go-playground/validator/v10"
	"hacktiv8-final-project4/config"
	"hacktiv8-final-project4/controller"
	"hacktiv8-final-project4/middleware"
	"hacktiv8-final-project4/repository"
	"hacktiv8-final-project4/service"
)

func CategoryRoutes(r *gin.Engine) {
	database := config.DatabaseConnection()
	categoryRepository := repository.NewCategoryRepository(database)
	validator := validatorV10.New()
	categoryService := service.NewCategoryService(categoryRepository, validator)
	categoryController := controller.NewCategoryController(*categoryService)
	router := r.Group("/categories")
	{
		router.Use(middleware.JwtMiddleware())
		router.GET("/", categoryController.GetCategories)
		router.POST("/", categoryController.CreateCategory)
		router.PATCH("/:id", categoryController.UpdateCategory)
		router.DELETE("/:id", categoryController.DeleteCategory)
	}
}
