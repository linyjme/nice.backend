package router

import (
	"asyncClient/api"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.POST("createApi", api.CreateApi) // 创建Api
	}
}
