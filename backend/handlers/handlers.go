package handlers

import (
    "net/http"
    "aura-print/models"

    "github.com/gin-gonic/gin"
)

var (
    products []models.Product
    news     []models.News
    idCounter = 1
)

func GetProducts(c *gin.Context) {
    c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product.ID = idCounter
    idCounter++
    products = append(products, product)

    c.JSON(http.StatusCreated, product)
}

func GetNews(c *gin.Context) {
    c.JSON(http.StatusOK, news)
}

func AddNews(c *gin.Context) {
    var newsItem models.News
    if err := c.ShouldBindJSON(&newsItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newsItem.ID = idCounter
    idCounter++
    news = append(news, newsItem)

    c.JSON(http.StatusCreated, newsItem)
}

func SubmitContact(c *gin.Context) {
    var contact models.ContactRequest
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Здесь можно добавить логику отправки email
    c.JSON(http.StatusOK, gin.H{"message": "Сообщение отправлено успешно"})
}