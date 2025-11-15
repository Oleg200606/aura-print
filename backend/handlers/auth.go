package handlers

import (
    "fmt"
    "net/http"
    "auraprint/database"
    "auraprint/models"
    "github.com/gin-gonic/gin"
)

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
    Message string `json:"message"`
    Admin   string `json:"admin"`
    Success bool   `json:"success"`
}

func Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        fmt.Printf("❌ JSON bind error: %s\n", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // Валидация входных данных
    if req.Username == "" || req.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
        return
    }

    fmt.Printf("🔐 Login attempt - Username: %s\n", req.Username)

    var admin models.Admin
    if err := database.DB.Where("username = ?", req.Username).First(&admin).Error; err != nil {
        fmt.Printf("❌ Admin not found: %s\n", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    fmt.Printf("✅ Found admin: %s\n", admin.Username)

    // Проверка пароля с использованием bcrypt
    if err := admin.CheckPassword(req.Password); err != nil {
        fmt.Printf("❌ Password mismatch: %s\n", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    fmt.Printf("🎉 Login successful for: %s\n", admin.Username)
    
    c.JSON(http.StatusOK, LoginResponse{
        Message: "Login successful",
        Admin:   admin.Username,
        Success: true,
    })
}