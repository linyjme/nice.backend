package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

//@author: yjLin
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func EncodeBase64(str string) string {
	strBytes := []byte(str)
	return base64.StdEncoding.EncodeToString(strBytes)
}

func DecodeBase64(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {

	}
	decodeStr := string(decoded)
	return decodeStr
}
