package routes

import "github.com/gin-gonic/gin"

func ProductRoutes(r *gin.Engine) {
	router := r.Group("/products")
	{
		router.POST("/add", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Add Product",
			})
		})
		router.PUT("/edit", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Edit Product",
			})
		})
		router.DELETE("/delete", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Delete Product",
			})
		})
		router.GET("/list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "List Product",
			})
		})
	}
}
