package util

import (
	"os"
	"path/filepath"
	"unsafe"
)

func GetProjectDirectory() string {
	dir, _ := os.Getwd()
	return dir
}

func GetConfigIniPath() string {
	projectDir := GetProjectDirectory()
	return filepath.Join(projectDir, "config", "config.yaml")
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
