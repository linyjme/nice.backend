package pkg

import (
	"os"
	"path"
	"unsafe"

	"niceBackend/common/global"

	"go.mongodb.org/mongo-driver/mongo"
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

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			Data string
			Cap  int
		}{s, len(s)},
	))
}
