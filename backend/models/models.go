package models

import (
    "time"
    "database/sql/driver"
    "encoding/json"
    "errors"
)

// User представляет администратора
type User struct {
    ID           int       `json:"id"`
    Username     string    `json:"username"`
    PasswordHash string    `json:"-"`
    Email        string    `json:"email"`
    Role         string    `json:"role"` // Добавьте это поле
    CreatedAt    time.Time `json:"created_at"`
}

// Product представляет товар
type Product struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    ImageURL    string    `json:"image_url"`
    Tags        Tags      `json:"tags"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// News представляет новость
type News struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    ImageURL  string    `json:"image_url"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// ContactRequest представляет запрос обратной связи
type ContactRequest struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
}

// Tags для хранения массива тегов в JSONB
type Tags []string

func (t Tags) Value() (driver.Value, error) {
    return json.Marshal(t)
}

func (t *Tags) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(b, &t)
}

// LoginRequest для аутентификации
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// AuthResponse ответ с JWT токеном
type AuthResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}