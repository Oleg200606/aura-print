package handlers

import (
    "net/http"
    "strconv"
    "aura-print/database"
    "aura-print/models"

    "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    rows, err := database.DB.Query(`
        SELECT id, name, description, image_url, tags, created_at, updated_at 
        FROM products 
        ORDER BY created_at DESC
    `)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        err := rows.Scan(
            &product.ID,
            &product.Name,
            &product.Description,
            &product.ImageURL,
            &product.Tags,
            &product.CreatedAt,
            &product.UpdatedAt,
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := database.DB.QueryRow(`
        INSERT INTO products (name, description, image_url, tags) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id, created_at, updated_at
    `, product.Name, product.Description, product.ImageURL, product.Tags).Scan(
        &product.ID, &product.CreatedAt, &product.UpdatedAt,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err = database.DB.Exec(`
        UPDATE products 
        SET name = $1, description = $2, image_url = $3, tags = $4, updated_at = CURRENT_TIMESTAMP 
        WHERE id = $5
    `, product.Name, product.Description, product.ImageURL, product.Tags, id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    _, err = database.DB.Exec("DELETE FROM products WHERE id = $1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}