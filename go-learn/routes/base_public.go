package routes

import (
	"go-learn/controller"

	"github.com/gin-gonic/gin"
)

// InitBaseRouter 初始化
func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.POST("/register", controller.Register)
	}
}
