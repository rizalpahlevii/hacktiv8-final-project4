package routes

import "github.com/gin-gonic/gin"

func TransactionRoutes(r *gin.Engine) {
	router := r.Group("/transactions")
	{
		router.GET("/user-transactions", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "List Transaction",
			})
		})
		router.GET("/my-transactions", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "List Transaction by User",
			})
		})
		router.POST("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Add Transaction",
			})
		})

	}
}
