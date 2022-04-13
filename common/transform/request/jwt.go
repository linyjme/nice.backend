package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	UUID        string
	ID          int64
	Account     string
	NickName    string
	AccountType uint8
	BufferTime  int64
	jwt.StandardClaims
}
