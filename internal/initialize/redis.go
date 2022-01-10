package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"niceBackend/common/global"
)

func Redis() {
	redisCfg := global.NiceConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.NiceLog.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.NiceLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.NiceRedis = client
	}
}
