package core

import (
	"niceBackend/common/global"
	"niceBackend/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	// 初始化redis服务
	initialize.Redis()
	//initialize.Mongo()
	address := fmt.Sprintf(":%d", global.NICE_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.NICE_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`欢迎使用 niceBackend`)
	global.NICE_LOG.Error(s.ListenAndServe().Error())
}
