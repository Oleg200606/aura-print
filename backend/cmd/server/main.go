package main

import (
	"auraprint/config"
	"auraprint/database"
	"auraprint/handlers"
	"auraprint/middleware"
	"os"
	"path/filepath"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.New()
	// Initialize database
	db, err := database.Connect(conf)
	if err != nil {
		log.Fatal("Failed to connect to database: %s", err.Error())
	}

	app := handlers.NewApplication(db)

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
		log.Printf("🌐 %s %s from %s\n", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr)
		c.Next()
	})

	// Public routes
	public := router.Group("/api")
	{
		public.GET("/products", app.GetProducts)
		public.GET("/news", app.GetNews)
		public.POST("/admin/login", app.Login)
		public.POST("/contact", app.SendContactMessage)

		// Новый маршрут для получения изображений
		public.GET("/images/:filename", app.GetImage)
	}

	// Admin routes
	admin := router.Group("/api/admin")
	{
		admin.POST("/products", app.CreateProduct)
		admin.DELETE("/products/:id", app.DeleteProduct)
		admin.POST("/news", app.CreateNews)
		admin.DELETE("/news/:id", app.DeleteNews)

		// Новые маршруты для загрузки изображений
		admin.POST("/upload/image", app.UploadImage)
		admin.DELETE("/images/:filename", app.DeleteImage)
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

	log.Print("🚀 Server starting on :8081")
	log.Print("📡 API available at http://localhost:8081/api")
	log.Print("❤️  Health check at http://localhost:8081/health")
	log.Print("🖼️  Images available at http://localhost:8081/uploads/")

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
