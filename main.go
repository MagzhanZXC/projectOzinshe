package main

import (
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

	r.Run(":8080") // Запуск сервера на порту 8080.
}
