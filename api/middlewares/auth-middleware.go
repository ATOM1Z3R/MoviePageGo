package middlewares

import (
	"api/consts"
	"api/utilities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No authorization header provided")})
			ctx.Abort()
			return
		}
		claims, err := utilities.ValidateToken(clientToken)
		if err != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
			return
		}

		if claims.TokenType != consts.TOKEN_ACCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong token type"})
			ctx.Abort()
			return
		}
		ctx.Set("userId", claims.UserId)
		ctx.Set("email", claims.Email)
		ctx.Set("firstName", claims.FirstName)
		ctx.Set("lastName", claims.LastName)
		ctx.Set("userRole", claims.UserRole)
		ctx.Set("tokenType", claims.TokenType)
		ctx.Next()
	}
}
