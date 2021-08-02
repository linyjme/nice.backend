package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"niceBackend/transform"
)

var (
	RAY_DB                  *gorm.DB
	RAY_CONFIG              transform.Server
	RAY_REDIS               *redis.Client
	RAY_VP                  *viper.Viper
	RAY_LOG                 *zap.Logger
	RAY_Mongo               *mongo.Client
	RAY_Concurrency_Control = &singleflight.Group{}
)
