package handlers

import (
    "net/http"
    "aura-print/database"
    "aura-print/models"

    "github.com/gin-gonic/gin"
)

func SubmitContact(c *gin.Context) {
    var contact models.ContactRequest
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := database.DB.QueryRow(`
        INSERT INTO contact_requests (name, email, phone, message) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id, created_at
    `, contact.Name, contact.Email, contact.Phone, contact.Message).Scan(
        &contact.ID, &contact.CreatedAt,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Contact request submitted successfully"})
}