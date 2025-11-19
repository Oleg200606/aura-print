package main

import (
	"auraprint/database"
	"auraprint/handlers"
	"auraprint/middleware"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Initialize database
	if err := database.InitDatabase(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	router := gin.Default()

	router.Use(middleware.Logger())
	router.Use(location.Default())

	// Настройка статических файлов
	setupStaticFiles(router)

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	}))

	// Log all requests
	router.Use(func(c *gin.Context) {
		fmt.Printf("🌐 %s %s from %s\n", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr)
		c.Next()
	})

	// Public routes
	public := router.Group("/api")
	{
		public.GET("/products", handlers.GetProducts)
		public.GET("/news", handlers.GetNews)
		public.POST("/admin/login", handlers.Login)
		public.POST("/contact", handlers.SendContactMessage)

		// Новый маршрут для получения изображений
		public.GET("/images/:filename", handlers.GetImage)
	}

	// Admin routes
	admin := router.Group("/api/admin")
	{
		admin.POST("/products", handlers.CreateProduct)
		admin.DELETE("/products/:id", handlers.DeleteProduct)
		admin.POST("/news", handlers.CreateNews)
		admin.DELETE("/news/:id", handlers.DeleteNews)

		// Новые маршруты для загрузки изображений
		admin.POST("/upload/image", handlers.UploadImage)
		admin.DELETE("/images/:filename", handlers.DeleteImage)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "AuraPrint Backend"})
	})

	// API info
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AuraPrint API",
			"version": "1.0",
			"endpoints": []string{
				"GET /api/products",
				"GET /api/news",
				"POST /api/admin/login",
				"POST /api/contact",
				"GET /api/images/:filename",
				"POST /api/admin/upload/image",
			},
		})
	})

	fmt.Println("🚀 Server starting on :8081")
	fmt.Println("📡 API available at http://localhost:8081/api")
	fmt.Println("❤️ Health check at http://localhost:8081/health")
	fmt.Println("🖼️ Images available at http://localhost:8081/uploads/")

	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed run server: %s", err)
	}
}

// setupStaticFiles настраивает обслуживание статических файлов
func setupStaticFiles(router *gin.Engine) {
	// Создаем папки для загрузок если их нет
	uploadDirs := []string{"uploads", "uploads/images", "uploads/products", "uploads/news"}
	for _, dir := range uploadDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Printf("⚠️  Failed to create directory %s: %v", dir, err)
		}
	}

	// Статические файлы для загрузок
	router.Static("/uploads", "./uploads")

	// Дополнительные статические пути
	router.Static("/static", "./static")
	router.Static("/assets", "./assets")

	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	// Fallback для изображений
	router.NoRoute(func(c *gin.Context) {
		urlIsImagePath := slices.Contains(imageExtensions, filepath.Ext(c.Request.URL.Path))

		if urlIsImagePath {
			c.File("./uploads" + c.Request.URL.Path)
			return
		}
		c.JSON(404, gin.H{"error": "Endpoint not found"})
	})
}
