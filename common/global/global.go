package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"niceBackend/common/transform"
)

var (
	NICE_DB                  *gorm.DB
	NICE_CONFIG              transform.Server
	NICE_VP                  *viper.Viper
	NICE_REDIS               *redis.Client
	NICE_LOG                 *zap.Logger
	NICE_Mongo               *mongo.Client
	AsyncChan                *transform.AsyncChan
	NICE_Concurrency_Control = &singleflight.Group{}
)
