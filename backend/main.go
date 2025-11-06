package main

import (
    "log"
    "os"
    "aura-print/database"
    "aura-print/handlers"
    "aura-print/middleware"

    "github.com/gin-gonic/gin"
)

func main() {

  
    
    // Инициализация базы данных
    dbConfig := database.Config{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     getEnv("DB_PORT", "5432"),
        User:     getEnv("DB_USER", "postgres"),
        Password: getEnv("DB_PASSWORD", "1"),
        DBName:   getEnv("DB_NAME", "AuraPrints"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }

    if err := database.Connect(dbConfig); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Установка JWT секрета
    if os.Getenv("JWT_SECRET") == "" {
        os.Setenv("JWT_SECRET", "your-secret-key-change-in-production")
    }

    router := gin.Default()
    // main.go после database.Connect()
    err := database.CreateAdminUser("admin", "admin123", "admin@auraprint.com")
    if err != nil {
        log.Printf("Admin user creation: %v", err)
    } else {
        log.Println("Admin user created/verified")
    }

    // CORS middleware
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

    // Публичные маршруты
    public := router.Group("/api")
    {
        public.GET("/products", handlers.GetProducts)
        public.GET("/news", handlers.GetNews)
        public.POST("/contact", handlers.SubmitContact)
        public.POST("/login", handlers.Login)
    }

    // Защищенные маршруты (требуют аутентификации)
    protected := router.Group("/api/admin")
    protected.Use(middleware.AuthMiddleware())
    protected.Use(middleware.AdminMiddleware())
    {
        // Продукты CRUD
        protected.POST("/products", handlers.CreateProduct)
        protected.PUT("/products/:id", handlers.UpdateProduct)
        protected.DELETE("/products/:id", handlers.DeleteProduct)

        // Новости CRUD
        protected.POST("/news", handlers.CreateNews)
        protected.PUT("/news/:id", handlers.UpdateNews)
        protected.DELETE("/news/:id", handlers.DeleteNews)
    }

    port := getEnv("PORT", "8080")
    log.Printf("Server starting on port %s", port)
    router.Run(":" + port)
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}