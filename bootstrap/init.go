package bootstrap

import (
	"niceBackend/common/global"
	"niceBackend/config"
	"niceBackend/internal/pkg/log"
)

func Init(configPath string) {
	err := config.Init(configPath)
	if err != nil {
		panic("init fail")
	}
	global.NiceLog = log.Init(config.GetConf().Zap)

}
