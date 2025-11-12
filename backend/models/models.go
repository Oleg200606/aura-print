package models

import (
    "time"
)

type Product struct {
    ID          uint      `json:"id" gorm:"primary_key"`
    Name        string    `json:"name" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Image       string    `json:"image_url"`
    Category    string    `json:"category"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type News struct {
    ID          uint      `json:"id" gorm:"primary_key"`
    Title       string    `json:"title" binding:"required"`
    Content     string    `json:"content" binding:"required"`
    Image       string    `json:"image_url"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Admin struct {
    ID       uint   `json:"id" gorm:"primary_key"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
}