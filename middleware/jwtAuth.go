package middleware

import (
	"diary_api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context){
		err := helper.ValidateJWT(context)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error":"Authentication Required"})
			context.Abort()
			return 
		}
		context.Next()
	}
}