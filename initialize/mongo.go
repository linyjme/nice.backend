package initialize

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"niceBackend/common/global"
)

func Mongo() {
	mongoCfg := global.RAY_CONFIG.Mongo
	var err error
	clientOptions := options.Client().ApplyURI(mongoCfg.Addr)
	// 连接到MongoDB
	global.RAY_Mongo, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		global.RAY_LOG.Error("mongo connect ping failed, err:", zap.Any("err", err))
	}
	// 检查连接
	err = global.RAY_Mongo.Ping(context.TODO(), nil)
	if err != nil {
		global.RAY_LOG.Error("mongo connect ping failed, err:", zap.Any("err", err))
	}
}
