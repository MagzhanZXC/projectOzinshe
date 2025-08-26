package main

import (
	"projectOzinshe/config"
	"projectOzinshe/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabse() // Подключаемся к базе данных
	r := gin.Default()      // Создаем новый экземпляр Gin

	r.GET("/ping", handlers.PingHandler)          // Регистрация обработчика для /ping
	r.GET("/movies", handlers.MoviesHandler)      // Регистрация обработчика для получения списка фильмов
	r.GET("/movies/:id", handlers.GetMovieByID)   // Регистрация обработчика для получения фильма по ID
	r.POST("/movies", handlers.CreateMovie)       // Регистрация обработчика для создания нового фильма
	r.PUT("/movies/:id", handlers.UpdateMovie)    // Регистрация обработчика для обновления фильма
	r.DELETE("/movies/:id", handlers.DeleteMovie) // Регистрация обработчика для удаления фильма

	r.Run(":8080") // Запуск сервера на порту 8080

}
