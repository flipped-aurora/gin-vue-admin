package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/temp"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Temp    temp.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
