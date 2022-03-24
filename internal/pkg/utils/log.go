package utils

import (
	"go.uber.org/zap"
	"niceBackend/common/global"
)

func catchErrInfo(logInfo string, err error) {
	if err != nil {
		global.NiceLog.Error(logInfo, zap.Any("err", err))
	}
}
