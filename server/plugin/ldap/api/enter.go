package api

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	LdapApi
}

var ApiGroupApp = new(ApiGroup)

var (
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
