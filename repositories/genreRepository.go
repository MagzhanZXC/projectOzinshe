package repositories

import (
	"context"
	"errors"
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

func (r *GenresRepository) FindById(c context.Context, id int) (models.Genre, error) { // получение жанра по id
	genre, ok := r.db[id]
	if !ok {
		return models.Genre{}, errors.New("genre not found")
	}

	return genre, nil
}

func (r *GenresRepository) FindAll(c context.Context) []models.Genre { // получить все жанры
	genres := make([]models.Genre, 0, len(r.db))
	for _, v := range r.db {
		genres = append(genres, v)
	}

	return genres
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

func (r *GenresRepository) Create(c context.Context, genre models.Genre) int {
	id := len(r.db) + 1
	genre.Id = id
	r.db[id] = genre
	return id
}

func (r *GenresRepository) Update(c context.Context, genre models.Genre) error {
	original := r.db[genre.Id]
	original.Title = genre.Title
	r.db[genre.Id] = original
	return nil
}
func (r *GenresRepository) Delete(c context.Context, id int) error {
	delete(r.db, id)
	return nil
}

// создание, обновление, удаление жанров
