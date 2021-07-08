package routes

import (
	"go-learn/controller"

	"github.com/gin-gonic/gin"
)

// InitMovieRouter 初始化
func InitMovieRouter(Router *gin.RouterGroup) {
	MovieRouter := Router.Group("movies")
	{
		MovieRouter.GET("/", controller.GetMovies)
	}
}
