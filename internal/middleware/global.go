package middleware

import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/global"
)

// 全局中间件定义中间
func GlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//t := time.Now()
		global.NiceLog.Info(c.Request.RequestURI)
		global.NiceLog.Info(c.Request.Method)
		// 设置变量到Context的key中，可以通过Get()取
		//c.Set("request", "中间件")
		global.NiceLog.Info("中间件执行完毕")
		//t2 := time.Since(t)
	}
}
