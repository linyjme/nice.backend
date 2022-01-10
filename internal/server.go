package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"niceBackend/common/global"
	"niceBackend/internal/cron/asynchronous"
	"niceBackend/internal/initialize"
	"niceBackend/pkg/shutdown"
	"time"
)

type server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func RunServer() {
	Router := initialize.Routers()
	// 初始化redis服务
	initialize.Redis()
	//initialize.Mongo()
	address := fmt.Sprintf(":%d", global.NiceConfig.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(100 * time.Microsecond)
	global.NiceLog.Info("server run success on ", zap.String("address", address))

	fmt.Println("欢迎使用 nice Backend")
	// 关闭
	go shutdown.NewHook().Close(
		func() {
			asynchronous.StopChan(global.AsyncChan)
		},
		func() {
			// 程序结束前关闭数据库链接
			db, err := global.NiceDb.DB()
			if err != nil{
				fmt.Println("关闭 db")
				return
			}
			err = db.Close()
			if err != nil{
				fmt.Println("关闭 db 异常")
				return
			}

		},
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			fmt.Println("关闭 nice Backend")
			if err := s.Shutdown(ctx); err != nil {
				global.NiceLog.Error("server shutdown err", zap.Error(err))
			}
		},
	)
	go asynchronous.Consumer(global.AsyncChan)
	global.NiceLog.Error(s.ListenAndServe().Error())
}
