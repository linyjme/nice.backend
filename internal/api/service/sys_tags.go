package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"niceBackend/internal/model"
	"niceBackend/pkg"
)

func InsertTags() () {
	var result []model.Tags
	database := pkg.GetDatabase()
	queryResult := database.Collection("tags").FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	fmt.Println(queryResult)
}
