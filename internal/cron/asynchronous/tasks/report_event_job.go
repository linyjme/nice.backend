package tasks

import (
	"fmt"

	"niceBackend/common/global"

	"go.uber.org/zap"
)

func ReportEventJob() {
	defer func() {
		if err := recover(); err != nil {
			global.NiceLog.Error("ReportEventJob run fail ", zap.String("ReportEventJob", fmt.Sprintf("error: %v", err)))
		}
	}()

}
