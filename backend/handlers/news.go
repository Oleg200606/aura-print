package handlers

import (
    "net/http"
    "strconv"
    "aura-print/database"
    "aura-print/models"

    "github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
    rows, err := database.DB.Query(`
        SELECT id, title, content, image_url, created_at, updated_at 
        FROM news 
        ORDER BY created_at DESC
    `)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var news []models.News
    for rows.Next() {
        var newsItem models.News
        err := rows.Scan(
            &newsItem.ID,
            &newsItem.Title,
            &newsItem.Content,
            &newsItem.ImageURL,
            &newsItem.CreatedAt,
            &newsItem.UpdatedAt,
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        news = append(news, newsItem)
    }

    c.JSON(http.StatusOK, news)
}

func CreateNews(c *gin.Context) {
    var newsItem models.News
    if err := c.ShouldBindJSON(&newsItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := database.DB.QueryRow(`
        INSERT INTO news (title, content, image_url) 
        VALUES ($1, $2, $3) 
        RETURNING id, created_at, updated_at
    `, newsItem.Title, newsItem.Content, newsItem.ImageURL).Scan(
        &newsItem.ID, &newsItem.CreatedAt, &newsItem.UpdatedAt,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newsItem)
}

func UpdateNews(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
        return
    }

    var newsItem models.News
    if err := c.ShouldBindJSON(&newsItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err = database.DB.Exec(`
        UPDATE news 
        SET title = $1, content = $2, image_url = $3, updated_at = CURRENT_TIMESTAMP 
        WHERE id = $4
    `, newsItem.Title, newsItem.Content, newsItem.ImageURL, id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News updated successfully"})
}

func DeleteNews(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
        return
    }

    _, err = database.DB.Exec("DELETE FROM news WHERE id = $1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}