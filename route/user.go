package route

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/middleware"
)

func UserRoutes(r *gin.Engine) {
	router := r.Group("/users")
	{
		router.Use(middleware.JwtMiddleware())
		router.POST("/topup", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Topup",
			})
		})

	}
}
