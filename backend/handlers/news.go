package handlers

import (
    "net/http"
    "auraprint/database"
    "auraprint/models"
    "github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
    var news []models.News
    if err := database.DB.Order("created_at desc").Find(&news).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, news)
}

func CreateNews(c *gin.Context) {
    var news models.News
    if err := c.ShouldBindJSON(&news); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&news).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, news)
}

func DeleteNews(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Where("id = ?", id).Delete(&models.News{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "News deleted"})
}