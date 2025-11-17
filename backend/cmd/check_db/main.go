// backend/check_db.go
package main

import (
    "auraprint/database"
    "auraprint/models"
    "fmt"
    "log"
)

func main() {
    if err := database.InitDatabase(); err != nil {
        log.Fatal(err)
    }

    var admins []models.Admin
    if err := database.DB.Find(&admins).Error; err != nil {
        log.Fatal(err)
    }

    fmt.Printf("👥 Found %d admin users:\n", len(admins))
    for i, admin := range admins {
        fmt.Printf("   %d: ID=%d, Username='%s', Password='%s'\n", 
            i+1, admin.ID, admin.Username, admin.Password)
    }

    // Также проверим другие таблицы
    var products []models.Product
    database.DB.Find(&products)
    fmt.Printf("📦 Found %d products\n", len(products))

    var news []models.News
    database.DB.Find(&news)
    fmt.Printf("📰 Found %d news items\n", len(news))
}