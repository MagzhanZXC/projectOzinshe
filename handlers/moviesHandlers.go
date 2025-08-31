package handlers

import (
	"net/http"
	"strconv"

	"projectOzinshe/config"
	"projectOzinshe/models"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) { // Функция для обработки запроса на /ping
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func MoviesHandler(c *gin.Context) { // Функция для обработки запроса на /movies
	var movies []models.Movie     // Создаем срез для хранения фильмов
	config.DB.Find(&movies)       // Получаем все фильмы из базы данных
	c.JSON(http.StatusOK, movies) // Отправляем список фильмов в формате JSON,
}

func GetMovieByID(c *gin.Context) { // Функция для получения фильма по ID
	id, err := strconv.Atoi(c.Param("id")) // Преобразуем параметр ID из строки в целое число
	if err != nil {                        // Если произошла ошибка при преобразовании, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}

	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie) // Отправляем найденный фильм в формате JSON

}

func CreateMovie(c *gin.Context) { // Функция для создания нового фильма
	// Проверяем, что тело запроса содержит корректные данные
	var newMovie models.Movie                           // Создаем переменную для хранения нового фильма
	if err := c.ShouldBindJSON(&newMovie); err != nil { // Если произошла ошибка при привязке JSON, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newMovie) // Создаем новый фильм в базе данных
	c.JSON(http.StatusCreated, newMovie)
}

func UpdateMovie(c *gin.Context) { // Функция для обновления фильма
	id, err := strconv.Atoi(c.Param("id")) // Преобразуем параметр ID из строки в целое число
	if err != nil {                        // Если произошла ошибка при преобразовании, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var movie models.Movie                                    // Создаем переменную для хранения фильма
	if err := config.DB.First(&movie, id).Error; err != nil { // Ищем фильм по ID в базе данных
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	if err := c.ShouldBindJSON(&movie); err != nil { // Проверяем, что тело запроса содержит корректные данные
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}

	config.DB.Save(&movie)       // Сохраняем обновленный фильм в базе данных
	c.JSON(http.StatusOK, movie) // Отправляем обновленный фильм в формате JSON
}

func DeleteMovie(c *gin.Context) { // Функция для удаления фильма
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	config.DB.Delete(&models.Movie{}, id) // Удаляем фильм из базы данных
	c.Status(http.StatusNoContent)        // Отправляем ответ с кодом 204 No Content
}
