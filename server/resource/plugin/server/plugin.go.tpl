package {{ .Package }}

import (
	"context"
	"{{.Module}}/plugin/{{ .Package }}/initialize"
	interfaces "{{.Module}}/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

// 如果需要配置文件，请到config.Config中填充配置结构，且到下方发放中填入其在config.yaml中的key并添加如下方法
// initialize.Viper()
// 安装插件时候自动注册的api数据请到下方法.Api方法中实现并添加如下方法
// initialize.Api(ctx)
// 安装插件时候自动注册的api数据请到下方法.Menu方法中实现并添加如下方法
// initialize.Menu(ctx)
func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background() 
	initialize.Gorm(ctx)
	initialize.Router(group)
}
