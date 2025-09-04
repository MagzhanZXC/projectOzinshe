package models

type Movie struct {
	Id          int
	Title       string
	Description string
	ReleaseYear int
	Director    string
	Rating      int
	IsWatched   bool
	TrailerURL  string
	PosterURL   string
	Genres      []Genre
}
