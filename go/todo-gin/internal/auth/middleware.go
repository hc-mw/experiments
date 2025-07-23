package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hardikchoksi151/todo-gin/pkg/jwt"
)

func Middleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ParseToken(tokenStr, secret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx.Set("userID", claims["user_id"])
		ctx.Next()
	}
}
