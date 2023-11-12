package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutes(r *gin.Engine) {
	router := r.Group("/swagger")
	{
		router.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
