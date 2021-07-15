package main

import (
	"fmt"
	"os"
	"path"
	"gopkg.in/ini.v1"
	"asyncClient/common/untils"
)

func main() {
	//dir,_ := os.Getwd()
	projectDir := untils.GetProjectDirectory()
	configIni := path.Join(projectDir, "conf", "config.ini")
	fmt.Println("当前路径：", projectDir)
	fmt.Println("ini路径：", configIni)
	cfg, err := ini.Load(configIni)
	if err != nil {
		fmt.Println("文件读取错误", err)
		os.Exit(1)
	}
	fmt.Println(cfg.Section("listen").Key("host"))
}
