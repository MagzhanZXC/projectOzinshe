package handlers

import (
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
		db: map[int]models.Movie{
			1: {
				ID:          1,
				Title:       "Вверх",
				Description: "Мультфильм о приключениях старика Карла и мальчика Рассела.",
				ReleaseYear: 2009,
				Director:    "Пит Доктер",
				Rating:      0,
				IsWatched:   false,
				TrailerURL:  "https://www.youtube.com/watch?v=ORFWdXl_zJ4",
				PosterURL:   "",
				Genre:       make([]models.Genre, 0),
			},
			2: {
				ID:          2,
				Title:       "Тачки",
				Description: "Фильм о гонках автомобилей и дружбе.",
				ReleaseYear: 2006,
				Director:    "Дэн Скэнлон",
				Rating:      0,
				IsWatched:   false,
				TrailerURL:  "https://www.youtube.com/watch?v=zSWdZVtXT7E",
				PosterURL:   "",
				Genre:       make([]models.Genre, 0),
			},
			3: {
				ID:          3,
				Title:       "Король Лев",
				Description: "Анимационный фильм о львенке Симбе и его приключениях.",
				ReleaseYear: 1994,
				Director:    "Роджер Аллерс",
				Rating:      0,
				IsWatched:   false,
				TrailerURL:  "https://www.youtube.com/watch?v=4sj1MT05lAA",
				PosterURL:   "",
				Genre:       make([]models.Genre, 0),
			},
		},
	}
}

func (h *MoviesHandler) FindAll(c *gin.Context) {
	movies := make([]models.Movie, 0, len(h.db))
	for _, movie := range h.db {
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MoviesHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid movie ID"))
		return
	}

	movie, ok := h.db[id]
	if !ok {
		c.JSON(http.StatusNotFound, models.NewApiError("Movie not found"))
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MoviesHandler) Create(c *gin.Context) {
	var m models.Movie
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind JSON"))
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
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid movie ID"))
		return
	}

	originalMovie, ok := h.db[id]
	if !ok {
		c.JSON(http.StatusNotFound, models.NewApiError("Movie not found"))
		return
	}

	var updateMovie models.Movie
	err = c.BindJSON(&updateMovie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind JSON"))
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

func (h *MoviesHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid movie ID"))
		return
	}

	delete(h.db, id)
	c.Status(http.StatusOK)
}
