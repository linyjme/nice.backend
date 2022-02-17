package tag_service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/repository/db_repo/tag_repo"
	"niceBackend/pkg"
)

func (s *service) Create(ctx core.Context) (id int32, err error) {
	var result []tag_repo.Tags
	database := pkg.GetDatabase()
	queryResult := database.Collection("tags").FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	fmt.Println(queryResult)
	return
}
