package global

import (
	"asyncClient/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	RAY_DB                  *gorm.DB
	RAY_CONFIG              common.Server
	RAY_VP                  *viper.Viper
	RAY_LOG                 *zap.Logger
	RAY_Concurrency_Control = &singleflight.Group{}
)
