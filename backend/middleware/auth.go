package middleware

import (
    "net/http"
    "strings"
    "aura-print/utils"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateToken(parts[1])
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("username", claims.Username)
        c.Set("userID", claims.UserID)
        c.Next()
    }
}

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Здесь можно добавить проверку ролей, если нужно
        c.Next()
    }
}