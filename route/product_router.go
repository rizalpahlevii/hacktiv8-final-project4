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

func ProductRoutes(r *gin.Engine) {
	database := config.DatabaseConnection()
	productRepository := repository.NewProductRepository(database)
	validator := validatorV10.New()
	productService := service.NewProductService(productRepository, validator)
	productController := controller.NewProductController(productService)

	router := r.Group("/products")
	{
		router.Use(middleware.JwtMiddleware())
		router.GET("/", productController.GetProducts)
		router.POST("/", productController.CreateProduct)
		router.PUT("/:id", productController.UpdateProduct)
		router.DELETE("/:id", productController.DeleteProduct)

	}
}
