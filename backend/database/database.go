package database

import (
	"auraprint/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open("sqlite3", "auraprint.db")
	if err != nil {
		return err
	}

	// Auto migrate models
	DB.AutoMigrate(&models.Product{}, &models.News{}, &models.Admin{})

	// Create default admin user with hashed password
	admin := models.Admin{
		Username: "admin",
	}

	// Хешируем пароль
	if err := admin.HashPassword("password"); err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Check if admin exists, if not create it
	var existingAdmin models.Admin
	if DB.Where("username = ?", "admin").First(&existingAdmin).Error != nil {
		if err := DB.Create(&admin).Error; err != nil {
			return err
		}
		fmt.Println("✅ Admin user created: admin / password")
	} else {
		fmt.Println("✅ Admin user already exists")
	}

	return nil
}
