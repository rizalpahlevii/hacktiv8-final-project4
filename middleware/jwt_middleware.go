package middleware

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/helper"
	"net/http"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is required.",
			})
			return
		}

		claims, err := helper.VerifyToken(token)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		} else {
			context.Set("claims", claims)
			context.Next()
		}
	}
}
