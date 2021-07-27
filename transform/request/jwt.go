package request

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Custom claims structure
type CustomClaims struct {
	ID          uint
	Account     string
	AccountType int
	CurrentTime time.Time
	BufferTime  int64
	jwt.StandardClaims
}
