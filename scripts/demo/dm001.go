package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func ByteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func DecodeBase64(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil{

	}
	decodeStr := string(decoded)
	return decodeStr
}

func main()  {
	str := "root"
	res := MD5V([]byte(str))
	fmt.Println(res)

	str2 := "cm9vdA=="
	res = DecodeBase64(str2)
	fmt.Println(res)

}
