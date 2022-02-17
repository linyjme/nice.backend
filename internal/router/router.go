package router

import (
	"niceBackend/common/global"
	"niceBackend/internal/alert"
	"niceBackend/internal/metrics"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/router/interceptor"
)

type resource struct {
	mux core.Mux
	interceptors interceptor.Interceptor
}

type Server struct {
	Mux core.Mux
}

func NewHTTPServer() (*Server, error) {
	r := new(resource)
	openBrowserUri := global.NiceConfig.System.Domain + string(global.NiceConfig.System.Port)

	mux, err := core.New(global.NiceLog,
		core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithAlertNotify(alert.NotifyHandler(global.NiceLog)),
		core.WithRecordMetrics(metrics.RecordHandler()),
	)

	if err != nil {
		panic(err)
	}
	r.mux = mux
	r.interceptors = interceptor.New(global.NiceLog)

	// 设置 Render 路由
	//setRenderRouter(r)

	// 设置 API 路由
	setApiRouter(r)

	// 设置 Socket 路由
	setSocketRouter(r)

	s := new(Server)
	s.Mux = mux
	return s, nil
}
