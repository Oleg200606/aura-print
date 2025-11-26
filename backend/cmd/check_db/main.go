// backend/check_db.go
package main

import (
	"auraprint/config"
	"auraprint/database"
	"auraprint/models"
	"fmt"

	"github.com/charmbracelet/log"
)

func main() {

	conf := config.New()

	db, err := database.Connect(conf)
	if err != nil {
		log.Fatal(err)
	}

	var admins []models.Admin
	if err := db.Find(&admins).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("👥 Found %d admin users:\n", len(admins))
	for i, admin := range admins {
		fmt.Printf("   %d: ID=%d, Username='%s', Password='%s'\n",
			i+1, admin.ID, admin.Username, admin.PasswordHash)
	}

	// Также проверим другие таблицы
	var products []models.Product
	db.Find(&products)
	fmt.Printf("📦 Found %d products\n", len(products))

	var news []models.News
	db.Find(&news)
	fmt.Printf("📰 Found %d news items\n", len(news))
}
