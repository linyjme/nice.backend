package initialize

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"niceBackend/common/global"
)

func Mongo() {
	mongoCfg := global.NiceConfig.Mongo
	clientOptions := options.Client().ApplyURI(mongoCfg.Addr)
	clientOptions.SetMaxPoolSize(mongoCfg.Number)
	// 设置连接池
	// 连接到MongoDB
	var err error
	global.NiceMongo, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		global.NiceLog.Error("mongo connect ping failed, err:", zap.Any("err", err))
	}
	// 检查连接
	err = global.NiceMongo.Ping(context.TODO(), nil)
	if err != nil {
		global.NiceLog.Error("mongo connect ping failed, err:", zap.Any("err", err))
	}
}
