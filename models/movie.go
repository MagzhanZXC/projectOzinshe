package models

type Movie struct {
	ID          int      `json:"id" gorm:"primary_key"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Year        int      `json:"year"`
	Director    string   `json:"director"`
	PosterURL   string   `json:"poster_url"`
	Genres      []string `json:"genres" gorm:"many2many:movie_genres;"`
	TrailerURL  string   `json:"trailer_url"`
}
