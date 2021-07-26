package router

import (
	"asyncClient/api/v1"
	"asyncClient/middleware"
	"github.com/gin-gonic/gin"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	EmailRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		EmailRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
