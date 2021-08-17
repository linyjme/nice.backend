package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/resources/api/announcement"
	"niceBackend/resources/api/auth"
	"niceBackend/resources/api/expansion"
)

func InitAdminRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.GET("expansion/statistic", expansion.GetStatistic)
		ApiRouter.POST("expansion/statistic", expansion.PostStatistic)
		ApiRouter.GET("auth/admin", auth.Admin)
		ApiRouter.GET("announcement", announcement.GetAnnouncement)
		ApiRouter.POST("announcement", announcement.PostAnnouncement)
	}
}
