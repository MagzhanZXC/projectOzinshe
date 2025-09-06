package handlers

import "projectOzinshe/models"

type GenresHandler struct {
	db map[int]models.Genre
}

func NewGenresHandler() *GenresHandler {
	return &GenresHandler{
		db: make(map[int]models.Genre),
	}
}

// реализовать методы для работы с жанрами
func (h *GenresHandler) FindAll(c) {
	// получить все жанры
}

func (h *GenresHandler) FindByID(c) {
	// получение жанра по id
}

func (h *GenresHandler) Create(c) {
	// создание жанра
}

func (h *GenresHandler) Update(c) {
	// обновление жанра
}

func (h *GenresHandler) Delete(c) {
	// удаление жанра
}
