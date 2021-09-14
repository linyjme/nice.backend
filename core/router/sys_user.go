package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/core/resources/api/user"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.POST("user/create", user.Register)
	}
}
