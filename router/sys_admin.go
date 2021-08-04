package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/resources/api/expansion"
)

func InitAdminRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.GET("expansion/statistic", expansion.Statistic)
	}
}

