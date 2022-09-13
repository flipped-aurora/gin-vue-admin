package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
