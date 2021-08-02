package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"niceBackend/common/global"
)

func Redis() {
	redisCfg := global.RAY_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.RAY_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.RAY_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.RAY_REDIS = client
	}
}
