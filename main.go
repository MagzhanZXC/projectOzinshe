package main

import (
	"context"
	"projectOzinshe/handlers"
	"projectOzinshe/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))

	_, err := connectToDb()
	if err != nil {
		panic(err)
	}

	moviesRepository := repositories.NewMoviesRepository()
	genresRepository := repositories.NewGenresRepository()
	moviesHandler := handlers.NewMoviesHandler(
		moviesRepository,
		genresRepository,
	)
	genresHandler := handlers.NewGenreHandlers(genresRepository)

	r.GET("/movies/:id", moviesHandler.FindById)
	r.GET("/movies", moviesHandler.FindAll)
	r.POST("/movies", moviesHandler.Create)
	r.PUT("/movies/:id", moviesHandler.Update)
	r.DELETE("/movies/:id", moviesHandler.Delete)

	r.GET("/genres/:id", genresHandler.FindById)
	r.GET("/genres", genresHandler.FindAll)
	r.POST("/genres", genresHandler.Create)
	r.PUT("/genres/:id", genresHandler.Update)
	r.DELETE("/genres/:id", genresHandler.Delete)

	r.Run(":8080")
}

func loadConfig() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil

}

func connectToDb() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
