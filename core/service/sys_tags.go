package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"niceBackend/core/model"
	"niceBackend/utils"
)

func InsertTags() () {
	var result []model.Tags
	database := utils.GetDatabase()
	queryResult := database.Collection("tags").FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	fmt.Println(queryResult)
}
