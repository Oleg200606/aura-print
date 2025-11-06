package database

import (
    "database/sql"
    "fmt"
    "log"
	_ "github.com/lib/pq"
    "aura-print/utils"

)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

var DB *sql.DB

func Connect(config Config) error {
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.User, config.Password, config.DBName,
    )
    
    // Добавьте эту строку для отладки
    log.Printf("Connection string: host=%s port=%s user=%s dbname=%s", 
        config.Host, config.Port, config.User, config.DBName)

    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Printf("Error opening database: %v", err)
        return err
    }

    if err = DB.Ping(); err != nil {
        log.Printf("Error pinging database: %v", err)
        return err
    }

    log.Println("Connected to PostgreSQL database")
    return nil
}

// database/database.go
func CreateAdminUser(username, password, email string) error {
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        return err
    }

    _, err = DB.Exec(`
        INSERT INTO users (username, password_hash, email) 
        VALUES ($1, $2, $3)
        ON CONFLICT (username) DO NOTHING
    `, username, hashedPassword, email)
    
    return err
}