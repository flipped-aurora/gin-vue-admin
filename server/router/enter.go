package router

import (
	"kirer.cn/server/router/example"
	"kirer.cn/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
