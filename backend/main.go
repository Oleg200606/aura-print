package main

import (
    "aura-print/handlers"

    "github.com/gin-gonic/gin"
    
)

func main() {
    router := gin.Default()

    // Настройка CORS
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })

    // Маршруты API
    api := router.Group("/api")
    {
        // Продукты
        api.GET("/products", handlers.GetProducts)
        api.POST("/products", handlers.AddProduct)
        
        // Новости
        api.GET("/news", handlers.GetNews)
        api.POST("/news", handlers.AddNews)
        
        // Контакты
        api.POST("/contact", handlers.SubmitContact)
    }

    // Статические файлы
    router.Static("/static", "./static")

    router.Run(":8080")
}