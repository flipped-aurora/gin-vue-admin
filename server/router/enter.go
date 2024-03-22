package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/home"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/webcms"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Webcms  webcms.RouterGroup
	Home    home.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
