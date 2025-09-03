package main

import (
	"projectOzinshe/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	corsConfig := cors.Config{ // Настройки CORS
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}

	r.Use(cors.New(corsConfig))

	moviesHandler := handlers.NewMoviesHandler() // Инициализация обработчика фильмов

	r.POST("/movies", moviesHandler.Create)    // Маршрут для создания фильма
	r.PUT("/movies/:id", moviesHandler.Update) // Маршрут для обновления фильма
	r.Run(":8080")                             // Запуск сервера на порту 8080.
}
