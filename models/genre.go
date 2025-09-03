package models

type Genre struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
}
