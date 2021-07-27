package core

import (
	"asyncClient/global"
	"asyncClient/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.RAY_CONFIG.Listen.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.RAY_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`欢迎使用 asyncClient`)
	global.RAY_LOG.Error(s.ListenAndServe().Error())
}
