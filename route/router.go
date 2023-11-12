package route

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/exception"
)

func StartApplication() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Hacktiv8 Final Project 4 API",
		})
	})

	gin.DefaultErrorWriter = exception.ErrorLogWriter
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	AuthRoutes(router)
	UserRoutes(router)
	CategoryRoutes(router)
	ProductRoutes(router)
	TransactionRoutes(router)
	SwaggerRoutes(router)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
