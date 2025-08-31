package handlers

import (
	"projectOzinshe/config"
	"projectOzinshe/models"

	"github.com/gin-gonic/gin"
)

func GenresHandler(c *gin.Context) {
	var genres []models.Genre
	config.DB.Find(&genres)
	c.JSON(200, genres)
}

func CreateGenre(c *gin.Context) {

}
