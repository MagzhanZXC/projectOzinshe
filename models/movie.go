package models

type Movie struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	PosterURL   string `json:"poster_url"`
}
