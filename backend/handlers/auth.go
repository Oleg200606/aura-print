package handlers

import (
    "net/http"
    "aura-print/database"
    "aura-print/models"
    "aura-print/utils"

    "github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
    var loginReq models.LoginRequest
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var user models.User
    err := database.DB.QueryRow(
        "SELECT id, username, password_hash, email FROM users WHERE username = $1",
        loginReq.Username,
    ).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email)

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if !utils.CheckPasswordHash(loginReq.Password, user.PasswordHash) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := utils.GenerateToken(user.ID, user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, models.AuthResponse{
        Token: token,
        User:  user,
    })
}