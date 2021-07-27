package utils

import "os"

func GetProjectDirectory() string {
	dir, _ := os.Getwd()
	return dir
}
