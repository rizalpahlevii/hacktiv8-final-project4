package route

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/exception"
)

func StartApplication() {
	router := gin.Default()
	gin.DefaultErrorWriter = exception.ErrorLogWriter
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	AuthRoutes(router)
	UserRoutes(router)
	CategoryRoutes(router)
	ProductRoutes(router)
	TransactionRoutes(router)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
