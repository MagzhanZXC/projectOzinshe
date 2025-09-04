package models

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseYear int
	Director    string
	Rating      int
	IsWatched   bool
	TrailerURL  string
	PosterURL   string
	Genre       []Genre
}
