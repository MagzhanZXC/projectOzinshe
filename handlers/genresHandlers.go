package handlers

import (
	"net/http"
	"projectOzinshe/config"
	"projectOzinshe/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GenresHandler(c *gin.Context) { // Функция для обработки запроса на /genres
	var genres []models.Genre     // Создаем срез для хранения жанров
	config.DB.Find(&genres)       // Получаем все жанры из базы данных
	c.JSON(http.StatusOK, genres) // Отправляем список жанров в формате JSON
}

func GetGenreByID(c *gin.Context) { // Функция для получения жанра по ID
	id, err := strconv.Atoi(c.Param("id")) // Преобразуем параметр ID из строки в целое число
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}

	var genre models.Genre
	if err := config.DB.First(&genre, id).Error; err != nil { // Ищем жанр по ID в базе данных
		c.JSON(http.StatusNotFound, gin.H{"error": "genre not found"}) // Отправляем сообщение об ошибке с кодом 404 Not Found
		return
	}
	c.JSON(http.StatusOK, genre) // Отправляем найденный жанр с кодом 200 OK
}

func CreateGenre(c *gin.Context) { // Функция для создания нового жанра
	var newGenre models.Genre                           // Создаем переменную для хранения нового жанра
	if err := c.ShouldBindJSON(&newGenre); err != nil { // Если произошла ошибка при привязке JSON, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}
	config.DB.Create(&newGenre)          // Создаем новый жанр в базе данных
	c.JSON(http.StatusCreated, newGenre) // Отправляем созданный жанр с кодом 201 Created
}

func UpdateGenre(c *gin.Context) { // Функция для обновления жанра
	id, err := strconv.Atoi(c.Param("id")) // Преобразуем параметр ID из строки в целое число
	if err != nil {                        // Если произошла ошибка при преобразовании, отправляем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var genre models.Genre                                    // Создаем переменную для хранения жанра
	if err := config.DB.First(&genre, id).Error; err != nil { // Ищем жвнр по ID в базе данных
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	if err := c.ShouldBindJSON(&genre); err != nil { // Проверяем, что тело запроса содержит корректные данные
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Отправляем сообщение об ошибке с кодом 400 Bad Request
		return
	}

	config.DB.Save(&genre)       // Сохраняем обновленный жанр в базе данных
	c.JSON(http.StatusOK, genre) // Отправляем обновленный жанр в формате JSON
}

func DeleteGenre(c *gin.Context) { // Функция для удаления жанра
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	config.DB.Delete(&models.Genre{}, id) // Удаляем жанр из базы данных
	c.Status(http.StatusNoContent)        // Отправляем ответ с кодом 204 No Content
}
