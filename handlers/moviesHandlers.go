package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"projectOzinshe/models"

	"github.com/gin-gonic/gin"
)

type MoviesHandler struct {
	db map[int]models.Movie
}

func NewMoviesHandler() *MoviesHandler {
	return &MoviesHandler{
		db: make(map[int]models.Movie),
	}
}

func (h *MoviesHandler) Create(c *gin.Context) {
	var m models.Movie
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err))
		return
	}
	id := len(h.db) + 1

	m.ID = id
	m.Genre = make([]models.Genre, 0) // Инициализация среза жанров

	h.db[id] = m

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *MoviesHandler) Update(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(errors.New("Invalid movie Id")))
		return
	}

	originalMovie, ok := h.db[id]
	if !ok {
		c.JSON(http.StatusNotFound, models.NewApiError(errors.New("Movie not found")))
		return
	}

	var updateMovie models.Movie
	err = c.BindJSON(&updateMovie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(errors.New("Could not bind JSON")))
		return
	}

	originalMovie.Title = updateMovie.Title
	originalMovie.Description = updateMovie.Description
	originalMovie.ReleaseYear = updateMovie.ReleaseYear
	originalMovie.Director = updateMovie.Director
	originalMovie.Rating = updateMovie.Rating
	originalMovie.IsWatched = updateMovie.IsWatched
	originalMovie.TrailerURL = updateMovie.TrailerURL
	//originalMovie.PosterURL = updateMovie.PosterURL
	//originalMovie.Genres = updateMovie.Genres

	h.db[id] = originalMovie

	c.Status(http.StatusOK)
}
