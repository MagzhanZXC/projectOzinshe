package repository

import (
	"context"
	"projectOzinshe/models"
)

type GenreRepository struct {
	db map[int]models.Genre
}

func NewGenreRepository() *GenreRepository {
	return &GenreRepository{
		db: make(map[int]models.Genre),
	}
}

func (r *GenreRepository) FindAllByIds(c context.Context, ids []int) ([]models.Genre, error) {
	genres := make([]models.Genre, 0, len(r.db))
	for _, v := range r.db {
		for _, id := range ids {
			if v.Id == id {
				genres = append(genres, v)
			}
		}
	}

	return genres, nil
}
