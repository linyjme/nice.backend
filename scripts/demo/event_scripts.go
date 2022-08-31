package main

import (
	"bufio"
	"fmt"
	"os"
)

//func MD5V(str []byte) string {
//	h := md5.New()
//	h.Write(str)
//	return hex.EncodeToString(h.Sum(nil))
//}
//
//func ByteString(p []byte) string {
//	for i := 0; i < len(p); i++ {
//		if p[i] == 0 {
//			return string(p[0:i])
//		}
//	}
//	return string(p)
//}
//
//func DecodeBase64(str string) string {
//	decoded, err := base64.StdEncoding.DecodeString(str)
//	if err != nil {
//
//	}
//	decodeStr := string(decoded)
//	return decodeStr
//}

func main() {
	args := os.Args
	argsLen := len(os.Args)
	filePath := "event.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 1; i < argsLen; i++ {
		oneParams := args[i] //获取输入的第一个参数
		write.WriteString(string(oneParams))
		write.WriteString("\n")
		fmt.Println(oneParams)
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
