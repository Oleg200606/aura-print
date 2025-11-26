package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type App struct {
	db     *gorm.DB
	config *viper.Viper
}

func NewApplication(db *gorm.DB, conf *viper.Viper) *App {
	return &App{db: db, config: conf}
}
