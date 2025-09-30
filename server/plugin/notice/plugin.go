package notice

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

// Register 注册插件
func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()
	// 初始化数据库
	initialize.Gorm(ctx)
	// 初始化路由
	initialize.Router(group)
	// 初始化菜单
	initialize.Menu(ctx)
	// 初始化 api
	initialize.Api(ctx)
}

// RouterPath 返回插件路由路径
func (p *plugin) RouterPath() string {
	return "notice"
}

// Name 返回插件名称
func (p *plugin) Name() string {
	return "notice"
}

// Description 返回插件描述
func (p *plugin) Description() string {
	return "通知中心插件，提供实时通知、在线用户管理等功能"
}

// Version 返回插件版本
func (p *plugin) Version() string {
	return "v1.0.0"
}
