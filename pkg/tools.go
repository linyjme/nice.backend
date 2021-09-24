package pkg

import (
	"go.mongodb.org/mongo-driver/mongo"
	"niceBackend/common/global"
	"os"
	"path"
)

func GetProjectDirectory() string {
	dir, _ := os.Getwd()
	return dir
}

func GetConfigIniPath() string {
	projectDir := GetProjectDirectory()
	return path.Join(projectDir, "config", "config.yaml")
}

func GetStaticAppPath() string {
	projectDir := GetProjectDirectory()
	return path.Join(projectDir, "dist", "app")
}

func GetDatabase() *mongo.Database {
	return global.NICE_Mongo.Database(global.NICE_CONFIG.Mongo.DB)
}

func getNextSequenceValue(sequenceName string) {
	//filter := bson.D{{"id", "Ash"}}
	//sequenceDocument:= GetDatabase().Collection(sequenceName).FindOneAndUpdate(context.TODO(), filter)
	//fmt.Println(sequenceDocument)
	
}
