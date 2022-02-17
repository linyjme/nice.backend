package initialize

import (
	"go.uber.org/zap"
	"niceBackend/common/global"
	nc "niceBackend/config"
	"niceBackend/pkg"
)

// @description   set system config,
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func SetSystemConfig(system nc.Server) {
	cs := pkg.StructToMap(system)
	for k, v := range cs {
		global.NiceVp.Set(k, v)
	}
	err := global.NiceVp.WriteConfig()
	if err != nil {
		global.NiceLog.Error("set SystemConfig  failed", zap.Any("err", err))
	}
}
