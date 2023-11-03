package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/model"
	"net/http"
)

func IsAdminMiddleware(context *gin.Context) {
	fmt.Println("Middleware")
	user := helper.GetLoggedUser(context)
	if user.Role != model.AdminRole {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized.",
		})
		return
	}
	context.Next()
}
