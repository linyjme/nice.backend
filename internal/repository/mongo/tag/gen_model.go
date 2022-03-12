package tag

import (
	ctx "context"
	"fmt"

	"niceBackend/common/global"
	"niceBackend/internal/repository/db_repo/tag_repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func NewModel() *tag_repo.Tags {
	return new(tag_repo.Tags)
}

func FindAll() error {
	database := global.NiceMongo.Database(global.NiceConfig.Mongo.DB)
	findOptions := options.Find()
	findOptions.SetLimit(1)
	cur, err := database.Collection("tags").Find(ctx.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		global.NiceLog.Error("查询失败", zap.Any("err", err))
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	var elem tag_repo.Tags
	for cur.Next(ctx.TODO()) {
		// create a value into which the single document can be decoded
		err := cur.Decode(&elem)
		if err != nil {
			global.NiceLog.Error(err.Error())
		}
	}

	if err := cur.Err(); err != nil {
		global.NiceLog.Error(err.Error())
	}
	// Close the cursor once finished
	cur.Close(ctx.TODO())
	fmt.Println(elem)
	return nil
}
