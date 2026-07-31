package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ClearToken(c *gin.Context) {
	setTokenCookie(c, "", -1, time.Unix(1, 0))
}

func SetToken(c *gin.Context, token string, maxAge int) {
	setTokenCookie(c, token, maxAge, time.Time{})
}

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-token")
	if token != "" {
		return token
	}
	token, _ = c.Cookie("x-token")
	return token
}

func setTokenCookie(c *gin.Context, value string, maxAge int, expires time.Time) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "x-token",
		Value:    value,
		Path:     "/",
		Expires:  expires,
		MaxAge:   maxAge,
		Secure:   requestUsesHTTPS(c.Request),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func requestUsesHTTPS(r *http.Request) bool {
	if r.TLS != nil {
		return true
	}
	forwardedProto := strings.SplitN(r.Header.Get("X-Forwarded-Proto"), ",", 2)[0]
	return strings.EqualFold(strings.TrimSpace(forwardedProto), "https")
}

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("system").Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Username
	}
}

// GetUserType 从Gin的Context中获取从jwt解析出来的用户类型
func GetUserType(c *gin.Context) system.UserType {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.UserType
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UserType
	}
}

func LoginToken(user system.Login) (token string, claims systemReq.CustomClaims, err error) {
	j := NewJWT()
	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.GetUUID(),
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: user.GetAuthorityId(),
		UserType:    user.GetUserType(),
	})
	token, err = j.CreateToken(claims)
	return
}

// LoginTokenWithExpire 签发登录 token 可携带 MustChangePwd 强制改密标记
func LoginTokenWithExpire(user system.Login, mustChangePwd bool) (token string, claims systemReq.CustomClaims, err error) {
	j := NewJWT()
	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.GetUUID(),
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: user.GetAuthorityId(),
		UserType:    user.GetUserType(),
	})
	claims.MustChangePwd = mustChangePwd
	token, err = j.CreateToken(claims)
	return
}
