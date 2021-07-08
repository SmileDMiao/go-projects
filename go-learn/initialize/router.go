package initialize

import (
	"go-learn/middleware"
	"go-learn/routes"

	"github.com/gin-gonic/gin"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	// https了
	// Router.Use(middleware.LoadTls())

	PublicGroup := Router.Group("")
	{
		routes.InitBaseRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())

	{
		routes.InitMovieRouter(PrivateGroup)
		routes.InitUserRouter(PrivateGroup)
		routes.InitArticleRouter(PrivateGroup)
	}
	return Router
}
