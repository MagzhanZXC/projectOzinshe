package handlers

import (
	"net/http"
	"strconv"

	"projectOzinshe/models"
	repositories "projectOzinshe/repositories"

	"github.com/gin-gonic/gin"
)

type MoviesHandler struct {
	db        map[int]models.Movie // to be removed
	genreRepo *repositories.GenresRepository
}

type createMovieRequest struct {
	Title       string
	Description string
	ReleaseYear int
	Director    string
	TrailerUrl  string
	GenreIds    []int
}

type updateMovieRequest struct {
	Title       string
	Description string
	ReleaseYear int
	Director    string
	TrailerUrl  string
	GenreIds    []int
}

func NewMoviesHandler(genreRepo *repositories.GenresRepository) *MoviesHandler {
	return &MoviesHandler{
		db: map[int]models.Movie{
			1: {
				Id:          1,
				Title:       "Вверх",
				Description: "Мультфильм о приключениях старика Карла и мальчика Рассела.",
				ReleaseYear: 2009,
				Director:    "Пит Доктер",
				Rating:      0,
				IsWatched:   false,
				TrailerURL:  "https://www.youtube.com/watch?v=ORFWdXl_zJ4",
				PosterURL:   "",
				Genres:      make([]models.Genre, 0),
			},
			2: {
				Id:          2,
				Title:       "Тачки",
				Description: "Фильм о гонках автомобилей и дружбе.",
				ReleaseYear: 2006,
				Director:    "Дэн Скэнлон",
				Rating:      0,
				IsWatched:   false,
				TrailerURL:  "https://www.youtube.com/watch?v=zSWdZVtXT7E",
				PosterURL:   "",
				Genres:      make([]models.Genre, 0),
			},
			3: {
				Id:          3,
				Title:       "Король Лев",
				Description: "Анимационный фильм о львенке Симбе и его приключениях.",
				ReleaseYear: 1994,
				Director:    "Роджер Аллерс",
				Rating:      0,
				IsWatched:   false,
				TrailerURL:  "https://www.youtube.com/watch?v=4sj1MT05lAA",
				PosterURL:   "",
				Genres:      make([]models.Genre, 0),
			},
		},
		genreRepo: genreRepo,
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
	var request createMovieRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind JSON"))
		return
	}

	genres := h.genreRepo.FindAllByIds(c, request.GenreIds)

	movie := models.Movie{
		Id:          len(h.db) + 1,
		Title:       request.Title,
		Description: request.Description,
		ReleaseYear: request.ReleaseYear,
		Director:    request.Director,
		TrailerURL:  request.TrailerUrl,
		Genres:      genres,
	}

	h.db[movie.Id] = movie

	c.JSON(http.StatusOK, gin.H{
		"id": movie.Id,
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

	var request updateMovieRequest
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind JSON"))
		return
	}

	genres := h.genreRepo.FindAllByIds(c, request.GenreIds)

	originalMovie.Title = request.Title
	originalMovie.Description = request.Description
	originalMovie.ReleaseYear = request.ReleaseYear
	originalMovie.Director = request.Director
	originalMovie.TrailerURL = request.TrailerUrl
	originalMovie.Genres = genres

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
