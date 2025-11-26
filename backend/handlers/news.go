package handlers

import (
	"auraprint/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) GetNews(c *gin.Context) {
	var news []models.News
	if err := app.db.Order("created_at desc").Find(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}

func (app *App) CreateNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.db.Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, news)
}

func (app *App) DeleteNews(c *gin.Context) {
	id := c.Param("id")
	if err := app.db.Where("id = ?", id).Delete(&models.News{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "News deleted"})
}
