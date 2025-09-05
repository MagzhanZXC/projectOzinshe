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
