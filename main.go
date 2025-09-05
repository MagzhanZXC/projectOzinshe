package main

import (
	"projectOzinshe/handlers"

	"projectOzinshe/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))

	moviesRepository := repositories.NewMoviesRepository()
	genresRepository := repositories.NewGenresRepository()
	moviesHandler := handlers.NewMoviesHandler(moviesRepository, genresRepository)

	r.GET("/movies/:id", moviesHandler.FindByID)
	r.GET("/movies", moviesHandler.FindAll)
	r.POST("/movies", moviesHandler.Create)
	r.PUT("/movies/:id", moviesHandler.Update)
	r.DELETE("/movies/:id", moviesHandler.Delete)

	r.Run(":8080")
}
