package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user" comment:"用户信息"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user" comment:"用户信息"`
	Token     string         `json:"token" comment:"token"`
	ExpiresAt int64          `json:"expiresAt" comment:"过期时间"`
}
