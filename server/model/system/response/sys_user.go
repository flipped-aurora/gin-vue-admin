package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user" comment:"用户信息"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user" comment:"用户信息"`
	Token     string         `json:"token" comment:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0NTg0NywidXNlcm5hbWUiOiLlvawiLCJhdmF0YXIiOiJodHRwczovL3d3dy5odW94aWFuLmNuL3VwbG9hZC9pbWFnZS8yMDIxLzcvMTUvYzg0ZDA1YjM1YWU1NThhNWUxMGYxNjNkZDg3ZjNmMmIuanBnIiwiZXhwIjoxNjQyMTMwODY2LCJpc3MiOiJhc3IiLCJuYmYiOjE2NDE1MjUwNjZ9.Huye4E5mUDG1KxQP7lhPVwve2yUx7RfGazerv6nW-uY"`
	ExpiresAt int64          `json:"expiresAt" comment:"过期时间" example:"1641791563"`
}
