package models

type Movie struct {
	ID          int     `json:"id" gorm:"primary_key"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseYear int     `json:"release_year"`
	Director    string  `json:"director"`
	Rating      int     `json:"rating"`
	IsWatched   bool    `json:"is_watched"`
	TrailerURL  string  `json:"trailer_url"`
	PosterURL   string  `json:"poster_url"`
	Genre       []Genre `json:"genres" gorm:"many2many:movie_genres;"`
}
