package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/FlyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/NestAirlinePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/UserTeemlinkPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System            system.RouterGroup
	Example           example.RouterGroup
	NestInfo          NestInfo.RouterGroup
	Nestrolepkg       Nestrolepkg.RouterGroup
	NestAirlinePkg    NestAirlinePkg.RouterGroup
	NestExecRecordPkg NestExecRecordPkg.RouterGroup
	FlyResultPkg      FlyResultPkg.RouterGroup
	UserTeemlinkPkg   UserTeemlinkPkg.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
