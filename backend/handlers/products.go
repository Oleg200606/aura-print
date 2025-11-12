package handlers

import (
    "net/http"
    "auraprint/database"
    "auraprint/models"
    "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    var products []models.Product
    if err := database.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}