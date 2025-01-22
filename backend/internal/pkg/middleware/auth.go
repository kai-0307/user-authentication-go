package middleware

import (
	"net/http"
	"strings"

	"github.com/cos-plat/backend/internal/pkg/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *auth.JWTService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
            return
        }

        // Bearer トークンの取り出し
        splitToken := strings.Split(authHeader, "Bearer ")
        if len(splitToken) != 2 {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "不正な認証形式です"})
            return
        }

        tokenStr := splitToken[1]
        claims, err := jwtService.ValidateToken(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "トークンが無効です"})
            return
        }

        // ユーザー情報をコンテキストに設定
        c.Set("user_id", claims.UserID)
        c.Set("username", claims.Username)

        c.Next()
    }
}
