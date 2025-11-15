// middleware/logger.go
package middleware

import (
    "fmt"
    "time"
    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        c.Next()
        
        duration := time.Since(start)
        fmt.Printf("🌐 %s %s %d %v\n", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
    }
}