package initialize

import (
	"niceBackend/common/global"
	"niceBackend/middleware"
	"niceBackend/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // https
	global.RAY_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 跨域
	global.RAY_LOG.Info("use middleware cors")
	global.RAY_LOG.Info("register swagger handler")
	Router.Static("/app", "./dist/app")
	//Router.StaticFS(AppStaticPath, http.Dir(AppStaticPath))
	//Router.StaticFile("/", "./resources/favicon.ico")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitApiRouter(PrivateGroup)                   // 注册功能api路由
		// Code generated by niceBackend Begin; DO NOT EDIT.
		// Code generated by niceBackend End; DO NOT EDIT.
	}
	global.RAY_LOG.Info("router register success")
	return Router
}
