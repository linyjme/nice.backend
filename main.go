package main

import (
	"niceBackend/common/global"
	"niceBackend/core"
	"niceBackend/core/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @contact.name linyj
// @contact.url http://localhost:8080/swagger/index.html
// @contact.email linyj@163.com
// @in header
// @name x-token
// @BasePath /
func main() {
	global.NICE_VP = initialize.Viper() // 初始化Viper
	global.NICE_LOG = initialize.Zap()  // 初始化zap日志库
	global.NICE_DB = initialize.Gorm()  // gorm连接数据库
	if global.NICE_DB != nil {
		initialize.SqlTables(global.NICE_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.NICE_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
