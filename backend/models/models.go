package models

import "time"

type Product struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Image       string   `json:"image"`
    Tags        []string `json:"tags"`
    CreatedAt   time.Time `json:"created_at"`
}

type News struct {
    ID      int       `json:"id"`
    Title   string    `json:"title"`
    Content string    `json:"content"`
    Image   string    `json:"image"`
    Date    string    `json:"date"`
    CreatedAt time.Time `json:"created_at"`
}

type ContactRequest struct {
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required"`
    Phone   string `json:"phone"`
    Message string `json:"message" binding:"required"`
}