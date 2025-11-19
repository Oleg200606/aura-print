package handlers

import (
	"fmt"

	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

// GetImage возвращает изображение
func GetImage(c *gin.Context) {
	filename := c.Param("filename")

	// Безопасная проверка имени файла
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filename"})
		return
	}

	// Пробуем разные папки
	possiblePaths := []string{
		"uploads/images/" + filename,
		"uploads/products/" + filename,
		"uploads/news/" + filename,
		"uploads/" + filename,
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			c.File(path)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
}

// UploadImage загружает изображение
func UploadImage(c *gin.Context) {
	// Получаем файл из формы
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No image file provided",
		})
		return
	}

	// Проверяем тип файла
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}

	openedFile, _ := file.Open()
	buffer := make([]byte, 512)
	_, err = openedFile.Read(buffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Cannot read file",
		})
		return
	}
	openedFile.Close()

	mimeType := http.DetectContentType(buffer)
	if !allowedTypes[mimeType] {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid file type. Only JPEG, PNG, GIF, and WebP are allowed",
		})
		return
	}

	// Генерируем уникальное имя файла
	ext := filepath.Ext(file.Filename)
	timestamp := time.Now().Format("20060102150405")
	newFilename := fmt.Sprintf("%s%s", timestamp, ext)

	// Определяем папку для сохранения
	folder := c.DefaultQuery("folder", "images")
	savePath := filepath.Join("uploads", folder, newFilename)

	// Создаем папку если не существует
	os.MkdirAll(filepath.Dir(savePath), 0755)

	// Сохраняем файл
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to save image: " + err.Error(),
		})
		return
	}

	// Возвращаем URL изображения
	imageURL := fmt.Sprintf("/%s/%s", folder, newFilename)
	url := location.Get(c)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Image uploaded successfully",
		"data": gin.H{
			"filename": newFilename,
			"url":      imageURL,
			"fullUrl":  url.Host + imageURL,
		},
	})
}

// DeleteImage удаляет изображение
func DeleteImage(c *gin.Context) {
	filename := c.Param("filename")

	if strings.Contains(filename, "..") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filename"})
		return
	}

	// Пробуем удалить из разных папок
	possiblePaths := []string{
		"uploads/images/" + filename,
		"uploads/products/" + filename,
		"uploads/news/" + filename,
		"uploads/" + filename,
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			if err := os.Remove(path); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Failed to delete image: " + err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Image deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "Image not found",
	})
}

// ServeImage обслуживает изображения с правильными заголовками
func ServeImage(c *gin.Context) {
	imagePath := c.Param("filepath")
	fullPath := filepath.Join("uploads", imagePath)

	// Проверяем существование файла
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// Устанавливаем правильные заголовки для кэширования
	c.Header("Cache-Control", "public, max-age=86400") // Кэш на 1 день
	c.File(fullPath)
}
