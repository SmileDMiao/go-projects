package routes

import (
	"go-learn/controller"

	"github.com/gin-gonic/gin"
)

// InitUserRouter 初始化
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.POST(":id/articles", controller.GetArticle)
	}
}
