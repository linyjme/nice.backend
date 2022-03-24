package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"niceBackend/internal/pkg/async"
)

var (
	NiceDb                 *gorm.DB
	NiceRedis              *redis.Client
	NiceLog                *zap.Logger
	AsyncChan              *async.AsyncChan
	NiceConcurrencyControl = &singleflight.Group{}
)
