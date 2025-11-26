package database

import (
	"auraprint/models"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func Connect(config *viper.Viper) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", config.GetString("database.connection_string"))
	if err != nil {
		return db, err
	}

	// Auto migrate models
	db.AutoMigrate(&models.Product{}, &models.News{}, &models.Admin{})

	// Create default admin user with hashed password
	admin := models.Admin{
		Username: "admin",
	}

	// Хешируем пароль
	if err := admin.SetPassword("password"); err != nil {
		return nil, fmt.Errorf("failed to set admin password: %v", err)
	}

	// Check if admin exists, if not create it
	var existingAdmin models.Admin
	if db.Where("username = ?", "admin").First(&existingAdmin).Error != nil {
		if err := db.Create(&admin).Error; err != nil {
			return nil, err
		}
		log.Info("✅ Admin user created: admin / password")
	} else {
		log.Info("✅ Admin user already exists")
	}

	return db, nil
}
