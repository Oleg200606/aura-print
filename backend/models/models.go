package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	*BaseModel
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Images      []Image `json:"images"`
	Category    string  `json:"category"`
}

type Image struct {
	*BaseModel
	URL string `json:"url"`
}

type News struct {
	*BaseModel
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image_url"`
}

type Admin struct {
	*BaseModel
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"-"`
}

// SetPassword хеширует пароль
func (a *Admin) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed hash password: %s", err)
	}
	a.PasswordHash = string(hashedPassword)
	return nil
}

// CheckPassword проверяет пароль
func (a *Admin) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(password))
}
