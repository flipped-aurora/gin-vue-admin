package ginx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID       uint
	Username string
	NickName string
	Dev      bool
}

func GetUserInfo(c *gin.Context) UserInfo {
	defaultUser := UserInfo{
		Username: "beiluo",
		ID:       3,
		NickName: "北洛",
		Dev:      true,
	}
	value, exists := c.Get("claims")
	if !exists {
		return defaultUser
	}
	claims, ok := value.(*request.BaseClaims)
	if !ok {
		return defaultUser
	}
	return UserInfo{Username: claims.Username, ID: claims.ID, NickName: claims.NickName}
}
