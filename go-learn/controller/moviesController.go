package controller

import (
	"go-learn/global"
	"go-learn/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMovies controller
func GetMovies(c *gin.Context) {
	var movies []model.Movie
	global.DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}
