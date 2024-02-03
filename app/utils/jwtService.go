package utils

import (
	"app/global"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyCustomClaims struct {
	UserId uint
	jwt.StandardClaims
}

func GiveJWT(id uint) (string, error) {
	// jwt相关配置
	var claims = MyCustomClaims{
		id,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*24*365*100 - 60,
			Issuer:    "yjddb",
			Subject:   "yjddb",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString([]byte(global.JWTKey))
	return ss, err
}
