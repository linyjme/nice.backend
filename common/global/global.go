package global

import (
	"niceBackend/common/transform"
	"niceBackend/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	NiceDb                 *gorm.DB
	NiceConfig             config.Server
	NiceVp                 *viper.Viper
	NiceRedis              *redis.Client
	NiceLog                *zap.Logger
	NiceMongo              *mongo.Client
	AsyncChan              *transform.AsyncChan
	NiceConcurrencyControl = &singleflight.Group{}
)
