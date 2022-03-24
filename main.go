package main

import (
	"flag"
	"niceBackend/bootstrap"
	"niceBackend/common/global"
	"niceBackend/config"
	"niceBackend/internal"
	"niceBackend/internal/pkg/async"
	"niceBackend/internal/pkg/cache"
	"niceBackend/internal/pkg/gorm"
	"niceBackend/pkg/util"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

var (
	confPath   string
	scriptName string
)

func init() {
	flag.StringVar(&confPath, "c", util.GetConfigIniPath(), "配置文件路径")
	flag.StringVar(&scriptName, "s", "", "运行内置数据库助手脚本")
	flag.Parse()
	bootstrap.Init(confPath)
}

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
	global.NiceDb = gorm.Gorm() // g orm连接数据库
	global.AsyncChan = async.Init()
	if global.NiceDb != nil {
		gorm.SqlTables(global.NiceDb) // 初始化表
	} else {
		panic(global.NiceDb)
	}
	// 初始化缓存服务
	cache.Init(config.GetConf().Redis)
	internal.RunServer()
}
