package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movies struct {
	ID          int    `json:"id"`
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
	// Отправляем список фильмов в формате JSON
	c.JSON(http.StatusOK, movies)
}

func getMovieByID(c *gin.Context) { // Функция для получения фильма по ID
	id, err := strconv.Atoi(c.Param("id")) // Преобразуем параметр ID из строки в целое число
	if err != nil {                        // Если произошла ошибка при преобразовании, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}

	for _, movie := range movies { // Перебираем список фильмов
		if movie.ID == id { // Если ID совпадает, отправляем фильм
			c.JSON(http.StatusOK, movie) // Отправляем фильм в формате JSON
			return                       // Возвращаем ответ с фильмом
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "movie not found"}) // Отправляем сообщение об ошибке, если фильм не найден
}

func createMovie(c *gin.Context) { // Функция для создания нового фильма
	// Проверяем, что тело запроса содержит корректные данные
	var newMovie Movies
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	c.JSON(http.StatusCreated, newMovie)
}

func updateMovie(c *gin.Context) { // Функция для обновления фильма
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedMovie Movies                                 // Исправлено на правильный тип
	if err := c.ShouldBindJSON(&updatedMovie); err != nil { // Проверяем, что тело запроса содержит корректные данные
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, m := range movies { // Обновляем цикл для поиска по ID
		if m.ID == id { // Если ID совпадает, обновляем фильм
			movies[i].Title = updatedMovie.Title
			movies[i].Description = updatedMovie.Description
			movies[i].Year = updatedMovie.Year
			movies[i].PosterURL = updatedMovie.PosterURL
			c.JSON(http.StatusOK, movies[i])
			return
		}
	}
}

func deleteMovie(c *gin.Context) { // Функция для удаления фильма
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, m := range movies {
		if m.ID == id {
			movies = append(movies[:i], movies[i+1:]...)             // Удаляем фильм из списка
			c.JSON(http.StatusOK, gin.H{"message": "movie deleted"}) // Отправляем сообщение об успешном удалении
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"}) // Отправляем сообщение об ошибке, если фильм не найден
}

func main() {
	r := gin.Default() // Создаем новый экземпляр Gin

	r.GET("/ping", pingHandler)          // Регистрация обработчика для /ping
	r.GET("/movies", moviesHandler)      // Регистрация обработчика для получения списка фильмов
	r.GET("/movies/:id", getMovieByID)   // Регистрация обработчика для получения фильма по ID
	r.POST("/movies", createMovie)       // Регистрация обработчика для создания нового фильма
	r.PUT("/movies/:id", updateMovie)    // Регистрация обработчика для обновления фильма
	r.DELETE("/movies/:id", deleteMovie) // Регистрация обработчика для удаления фильма

	r.Run(":8080") // Запуск сервера на порту 8080 // Запуск сервера

}
