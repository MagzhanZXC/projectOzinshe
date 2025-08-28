package config

import (
	"log"

	"projectOzinshe/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Глобальная переменная для базы данных

func ConnectDatabase() { // ✅ исправил название
	dsn := "host=localhost user=postgres password=1234 dbname=movies_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}
	DB = database

	// Автомиграция для таблицы movies
	DB.AutoMigrate(&models.Movie{})
}
