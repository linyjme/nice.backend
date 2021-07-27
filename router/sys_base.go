package router

import (
	"asyncClient/api/site"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) () {

	BaseRouter := Router.Group("api")
	{
		BaseRouter.GET("website/info", site.SiteInfoResource) // 创建Api
	}
}
