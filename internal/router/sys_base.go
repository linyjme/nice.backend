package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/internal/resources/api/auth"
)

func InitBaseRouter(Router *gin.RouterGroup) () {

	BaseRouter := Router.Group("api")
	{
		BaseRouter.POST("auth/login", auth.Login) // 创建Api
		//BaseRouter.GET("auth/admin", auth.Admin) // 创建Api
	}
}
