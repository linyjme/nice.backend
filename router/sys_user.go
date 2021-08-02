package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/resources/api/user"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.POST("login", user.Login)
		
	}
}
