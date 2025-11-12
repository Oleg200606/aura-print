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
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fmt.Printf("🔐 Login attempt - Username: %s, Password: %s\n", req.Username, req.Password)

    var admin models.Admin
    if err := database.DB.Where("username = ?", req.Username).First(&admin).Error; err != nil {
        fmt.Printf("❌ Admin not found: %s\n", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
        return
    }

    fmt.Printf("✅ Found admin: %s, Stored password: %s\n", admin.Username, admin.Password)

    // SIMPLE PASSWORD COMPARISON (for development)
    if admin.Password != req.Password {
        fmt.Printf("❌ Password mismatch: expected '%s', got '%s'\n", admin.Password, req.Password)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    fmt.Printf("🎉 Login successful for: %s\n", admin.Username)
    
    c.JSON(http.StatusOK, LoginResponse{
        Message: "Login successful",
        Admin:   admin.Username,
        Success: true,
    })
}