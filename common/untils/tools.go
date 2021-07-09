package untils

import "os"
import "fmt"

func GetProjectDirectory() string {
	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	return dir
}

func GetSqlConnectInfo()  string{


}