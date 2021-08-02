package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/resources/api/user"
)

func InitBaseRouter(Router *gin.RouterGroup) () {

	BaseRouter := Router.Group("api")
	{
		BaseRouter.GET("user/login", user.Login) // 创建Api
		BaseRouter.POST("user/create", user.Register)
	}
}
