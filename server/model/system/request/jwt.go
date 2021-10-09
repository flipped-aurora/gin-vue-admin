package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	UUID        string
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
