package service

import (
	"go.uber.org/zap"
	"niceBackend/common/global"
	"niceBackend/transform"
	"niceBackend/utils"
)

// @description   set system config,
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func SetSystemConfig(system transform.Server) {
	cs := utils.StructToMap(system)
	for k, v := range cs {
		global.NICE_VP.Set(k, v)
	}
	err := global.NICE_VP.WriteConfig()
	if err != nil {
		global.NICE_LOG.Error("set SystemConfig  failed", zap.Any("err", err))
	}
}
