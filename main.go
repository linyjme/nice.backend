package main

import (
	"asyncClient/core"
	"asyncClient/global"
	"asyncClient/initialize"
	"asyncClient/utils"
	"path"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	projectDir := utils.GetProjectDirectory()
	configIni := path.Join(projectDir, "config", "config.ini")
	global.RAY_VP = initialize.Viper(configIni) // 初始化Viper
	global.RAY_LOG = initialize.Zap()           // 初始化zap日志库
	global.RAY_DB = initialize.Gorm()           // gorm连接数据库
	if global.RAY_DB != nil {
		initialize.SqlTables(global.RAY_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.RAY_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
