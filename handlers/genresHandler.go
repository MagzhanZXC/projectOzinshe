package handlers

import (
	"net/http"
	"projectOzinshe/models"
	"projectOzinshe/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenreHandlers struct {
	repo *repositories.GenresRepository
}

func NewGenreHandlers(repo *repositories.GenresRepository) *GenreHandlers {
	return &GenreHandlers{
		repo: repo,
	}
}

// реализовать методы для работы с жанрами

func (h *GenreHandlers) FindById(c *gin.Context) {
	idstr := c.Param("id")         // получение id из параметров запроса
	id, err := strconv.Atoi(idstr) // преобразование id в int
	if err != nil {                // если ошибка преобразования
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid Genre Id")) // возвращаем ошибку
		return
	}
	genre, err := h.repo.FindById(c, id) // получение жанра по id из репозитория
	if err != nil {                      // если ошибка получения жанра
		c.JSON(http.StatusBadRequest, models.NewApiError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, genre)
	// получение жанра по id
}

func (h *GenreHandlers) FindAll(c *gin.Context) {
	genres := h.repo.FindAll(c)

	c.JSON(http.StatusOK, genres)
	// получить все жанры
}

func (h *GenreHandlers) Create(c *gin.Context) {
	var g models.Genre
	err := c.BindJSON(&g)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	id := h.repo.Create(c, g)

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	// создание жанра
}

func (h *GenreHandlers) Update(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid Genre Id"))
		return
	}

	_, err = h.repo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err.Error()))
		return
	}

	var updatedGenre models.Genre
	err = c.BindJSON(&updatedGenre)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.repo.Update(c, id, updatedGenre)

	c.Status(http.StatusOK)
	// обновление жанра
}

func (h *GenreHandlers) Delete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid Genre Id"))
		return
	}

	_, err = h.repo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err.Error()))
		return
	}

	h.repo.Delete(c, id)

	c.Status(http.StatusOK)
	// удаление жанра
}
