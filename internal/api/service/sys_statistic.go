package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"niceBackend/internal/model"
)
import "niceBackend/pkg"

func GetStatistic() () {
	var result []model.Tags
	database := pkg.GetDatabase()
	queryResult := database.Collection("tags").FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	fmt.Println(queryResult)
}


