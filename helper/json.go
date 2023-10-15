package helper

import (
	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(context *gin.Context, result interface{}) {
	err := context.ShouldBindJSON(result)
	PanicIfError(err)
}
