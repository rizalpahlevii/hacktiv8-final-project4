package route

import "github.com/gin-gonic/gin"

func CategoryRoutes(r *gin.Engine) {
	router := r.Group("/categories")
	{
		router.POST("/add", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Add Category",
			})
		})
		router.PUT("/edit", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Edit Category",
			})
		})
		router.DELETE("/delete", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Delete Category",
			})
		})
		router.GET("/list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "List Category",
			})
		})
	}
}
