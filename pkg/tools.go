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

func GetDatabase() *mongo.Database {
	return global.NiceMongo.Database(global.NiceConfig.Mongo.DB)
}
