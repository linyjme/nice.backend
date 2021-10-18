package tasks

import (

	"fmt"
	"go.uber.org/zap"
	"niceBackend/common/global"
)


func ReportEventJob() {
	defer func() {
		if err := recover(); err != nil {
			global.NICE_LOG.Error("ReportEventJob run fail ", zap.String("ReportEventJob", fmt.Sprintf("error: %v", err)))
		}
	}()

}
