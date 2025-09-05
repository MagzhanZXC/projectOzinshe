package repositories

import (
	"context"
	"projectOzinshe/models"
)

type GenresRepository struct {
	db map[int]models.Genre // to be removed

}

func NewGenresRepository() *GenresRepository {
	return &GenresRepository{
		db: map[int]models.Genre{
			1: {Id: 1, Title: "Комедия"},
			2: {Id: 2, Title: "Мультфильм"},
			3: {Id: 3, Title: "Приключения"},
		},
	}
}

func (r *GenresRepository) FindAllByIds(c context.Context, ids []int) []models.Genre {
	genres := make([]models.Genre, 0, len(r.db))
	for _, v := range r.db {
		for _, id := range ids {
			if v.Id == id {
				genres = append(genres, v)
			}
		}
	}

	return genres
}
