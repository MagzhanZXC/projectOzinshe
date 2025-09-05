package repositories

import (
	"context"
	"errors"
	"projectOzinshe/models"
)

type MoviesRepository struct {
	db map[int]models.Movie // to be removed
}

func NewMoviesRepository() *MoviesRepository {
	return &MoviesRepository{
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
	}
}

func (r *MoviesRepository) FindById(c context.Context, id int) (models.Movie, error) {
	movie, ok := r.db[id]
	if !ok {
		return models.Movie{}, errors.New("Movie not found")
	}
	return movie, nil
}

func (r *MoviesRepository) FindAll(c context.Context) []models.Movie {
	movies := make([]models.Movie, 0, len(r.db))
	for _, movie := range r.db {
		movies = append(movies, movie)
	}
	return movies
}
