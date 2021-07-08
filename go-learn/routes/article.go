package routes

import (
	"go-learn/controller"

	"github.com/gin-gonic/gin"
)

// InitArticleRouter 初始化
func InitArticleRouter(Router *gin.RouterGroup) {
	ArticleRouter := Router.Group("articles")
	{
		ArticleRouter.POST("/", controller.CreateArticle)
		ArticleRouter.GET(":id", controller.GetArticle)
	}
}
