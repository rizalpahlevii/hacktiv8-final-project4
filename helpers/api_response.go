package helpers

import "github.com/gin-gonic/gin"

func JsonResponse(context *gin.Context, code int, data interface{}) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, data)
}
