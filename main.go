package main

import (
	"niceBackend/common/global"
	"niceBackend/internal"
	"niceBackend/internal/initialize"
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
	global.NiceVp = initialize.Viper() // 初始化Viper
	global.NiceLog = initialize.Zap()  // 初始化zap日志库
	global.NiceDb = initialize.Gorm()  // gorm连接数据库
	global.AsyncChan = initialize.InitAsync()
	if global.NiceDb != nil {
		initialize.SqlTables(global.NiceDb) // 初始化表
	}
	internal.RunServer()
}
