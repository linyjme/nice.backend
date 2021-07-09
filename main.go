package main

//import (
//	_ "raysyncClient/routers"
//	beego "github.com/beego/beego/v2/server/web"
//)
//
//func main() {
//	beego.Run()
//}

//type person struct {
//	name string
//	city string
//	age int8
//}
//
//func main()  {
//	var p1 person
//	//p1.age = 18
//	//p1.name = "mine"
//	//p1.city = "shenzhen"
//	fmt.Printf("p1=%v\n", p1)
//	fmt.Printf("p1=%#v\n", p1)
//
//}

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path"
	"raysyncClient/common/untils"
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
