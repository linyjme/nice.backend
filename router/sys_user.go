package router

import (
	"niceBackend/api/user"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.POST("login", user.Login)
		
	}
}
