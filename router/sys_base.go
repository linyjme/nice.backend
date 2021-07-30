package router

import (
	"niceBackend/api/site"
	"niceBackend/api/user"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) () {

	BaseRouter := Router.Group("api")
	{
		BaseRouter.GET("website/info", site.SiteInfoResource) // 创建Api
		BaseRouter.GET("user/login", user.Login) // 创建Api
		BaseRouter.POST("user/create", user.Register)
	}
}
