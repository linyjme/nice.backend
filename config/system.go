package config

import (
	"niceBackend/common/global"
	"niceBackend/pkg/util"

	"go.uber.org/zap"
)

// @description   set system config,
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func SetSystemConfig(system System) {
	cs := util.StructToMap(system)
	for k, v := range cs {
		GetVip().Set(k, v)
	}
	err := GetVip().WriteConfig()
	if err != nil {
		global.NiceLog.Error("set SystemConfig  failed", zap.Any("err", err))
	}
}
