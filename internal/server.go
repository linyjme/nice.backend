package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"niceBackend/common/global"
	"niceBackend/config"
	"niceBackend/internal/cron/asynchronous"
	"niceBackend/internal/router"
	"niceBackend/pkg/shutdown"
	"time"
)

func RunServer() {
	s, err := router.NewHTTPServer()
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf(":%d", config.GetConf().System.Port)
	global.NiceLog.Info("server run success on ", zap.String("address", address))
	server := &http.Server{
		Addr:           address,
		Handler:        s.Mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
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
			if err != nil {
				fmt.Println("关闭 db")
				return
			}
			err = db.Close()
			if err != nil {
				fmt.Println("关闭 db 异常")
				return
			}

		},
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			fmt.Println("关闭 nice Backend")
			if err := server.Shutdown(ctx); err != nil {
				global.NiceLog.Error("server shutdown err", zap.Error(err))
			}
		},
	)
	go asynchronous.Consumer(global.AsyncChan)
	global.NiceLog.Error(server.ListenAndServe().Error())
}
