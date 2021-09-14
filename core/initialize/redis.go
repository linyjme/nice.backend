package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"niceBackend/common/global"
)

func Redis() {
	redisCfg := global.NICE_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.NICE_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.NICE_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.NICE_REDIS = client
	}
}
