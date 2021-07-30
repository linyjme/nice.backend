package utils

import (
	"os"
	"path"
)

func GetProjectDirectory() string {
	dir, _ := os.Getwd()
	return dir
}

func GetConfigIniPath() string {
	projectDir := GetProjectDirectory()
	return path.Join(projectDir, "config", "config.ini")
}

func GetStaticAppPath() string {
	projectDir := GetProjectDirectory()
	return path.Join(projectDir, "dist", "app")
}