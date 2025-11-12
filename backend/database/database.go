package database

import (
    "auraprint/models"
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
    
    // Create default admin user with SIMPLE PASSWORD (for development)
    admin := models.Admin{
        Username: "admin",
        Password: "password", // plain text password
    }
    
    // Check if admin exists, if not create it
    var existingAdmin models.Admin
    if DB.Where("username = ?", "admin").First(&existingAdmin).Error != nil {
        if err := DB.Create(&admin).Error; err != nil {
            return err
        }
        println("✅ Admin user created: admin / password")
    } else {
        println("✅ Admin user already exists: admin / password")
    }
    
    return nil
}