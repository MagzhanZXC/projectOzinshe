package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movies struct {
	ID          int    `json:"id qorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	PosterURL   string `json:"poster_url"`
}

var db *qorm.DB // Глобальная переменная для базы данных

func pingHandler(c *gin.Context) { // Функция для обработки запроса на /ping
	// Отправляем ответ с сообщением "pong"
	// Используем c.JSON для отправки JSON-ответа
	// http.StatusOK - это код состояния HTTP 200 OK
	// gin.H - это сокращение для создания карты (map) с ключами и значениями
	// В данном случае мы отправляем сообщение "pong"
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func moviesHandler(c *gin.Context) { // Функция для обработки запроса на /movies
	var movies []Movies			   // Создаем срез для хранения фильмов
	db.Find(&movies)			   // Получаем все фильмы из базы данных
	c.JSON(http.StatusOK, movies) // Отправляем список фильмов в формате JSON
}

func getMovieByID(c *gin.Context) { // Функция для получения фильма по ID
	id, err := strconv.Atoi(c.Param("id")) // Преобразуем параметр ID из строки в целое число
	if err != nil {                        // Если произошла ошибка при преобразовании, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}

	var movie Movies
	if err := db.First(&movie, id).Error; err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)	// Отправляем найденный фильм в формате JSON

}


func createMovie(c *gin.Context) { // Функция для создания нового фильма
	// Проверяем, что тело запроса содержит корректные данные
	var newMovie Movies
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&newMovie) 			 // Создаем новый фильм в базе данных
	c.JSON(http.StatusCreated, newMovie) 
}

func updateMovie(c *gin.Context) { // Функция для обновления фильма
	id, err := strconv.Atoi(c.Param("id")) // Преобразуем параметр ID из строки в целое число
	if err != nil { 					  // Если произошла ошибка при преобразовании, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var movie Movies // Создаем переменную для хранения фильма
	if err := db.First(&movie, id).Error; err != nil { // Ищем фильм по ID в базе данных
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	if err := c.ShouldBindJSON(&movie); err != nil { // Проверяем, что тело запроса содержит корректные данные
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}

	db.Save(&movie) 			 // Сохраняем обновленный фильм в базе данных
	c.JSON(http.StatusOK, movie) // Отправляем обновленный фильм в формате JSON
}


func deleteMovie(c *gin.Context) { // Функция для удаления фильма
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	db.Delete(&Movies{}, id) // Удаляем фильм из базы данных
	c.JSON(http.StatusNoContent) // Отправляем ответ с кодом 204 No Content
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=movies_db port=5432 sslmode=disable"
	var err error
	db, err = qorm.Open(postgres.Open(dsn), &qorm.Config{}
	if err != nil {
		log.Fatal("failed to connect database:", err)	
	})
	
	db.AutoMigrate(&Movies{}) // Автоматически создаем таблицу Movies в базе данных
	r := gin.Default() // Создаем новый экземпляр Gin

	r.GET("/ping", pingHandler)          // Регистрация обработчика для /ping
	r.GET("/movies", moviesHandler)      // Регистрация обработчика для получения списка фильмов
	r.GET("/movies/:id", getMovieByID)   // Регистрация обработчика для получения фильма по ID
	r.POST("/movies", createMovie)       // Регистрация обработчика для создания нового фильма
	r.PUT("/movies/:id", updateMovie)    // Регистрация обработчика для обновления фильма
	r.DELETE("/movies/:id", deleteMovie) // Регистрация обработчика для удаления фильма

	r.Run(":8080") // Запуск сервера на порту 8080 // Запуск сервера

}
