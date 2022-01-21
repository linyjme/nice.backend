package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"niceBackend/internal/api/repository/db_repo/tag_repo"
)
import "niceBackend/pkg"

func GetStatistic() () {
	var result []tag_repo.Tags
	database := pkg.GetDatabase()
	queryResult := database.Collection("tags").FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	fmt.Println(queryResult)
}


