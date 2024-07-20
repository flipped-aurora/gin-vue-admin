package announcement

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()
	// 如果需要配置文件，请到config.Config中填充配置结构，且到下方发放中填入其在config.yaml中的key
	// initialize.Viper()
	// 安装插件时候自动注册的api数据请到下方法.Api方法中实现
	initialize.Api(ctx)
	// 安装插件时候自动注册的api数据请到下方法.Menu方法中实现
	initialize.Menu(ctx)
	initialize.Gorm(ctx)
	initialize.Router(group)
}
