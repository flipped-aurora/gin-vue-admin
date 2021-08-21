package router

import (
	"github.com/flipped-aurora/gin-vue-admin/router/autocode"
	"github.com/flipped-aurora/gin-vue-admin/router/example"
	"github.com/flipped-aurora/gin-vue-admin/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
