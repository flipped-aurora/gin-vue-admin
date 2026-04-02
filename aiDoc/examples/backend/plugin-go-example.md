# plugin.go 示例

## 这个文件负责什么

插件入口负责把插件注册进系统，并在系统启动时完成路由、菜单、字典、表结构等初始化。

## 什么时候应该这样写

- 新增一个标准 gin-vue-admin 插件
- 需要接入插件注册机制
- 需要把插件初始化逻辑收敛到统一入口

## 推荐写法示例

```go
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

func init() {
	interfaces.Register(Plugin)
}

func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()
	initialize.Api(ctx)
	initialize.Menu(ctx)
	initialize.Dictionary(ctx)
	initialize.Gorm(ctx)
	initialize.Router(group)
}
```

## 为什么这样写

- `init()` 中注册插件，系统启动时就能发现它
- `Register` 里把插件初始化步骤按顺序集中起来
- 插件入口文件保持薄，真正细节下沉到 `initialize/`

## 常见错误

- 在 `plugin.go` 里塞过多业务细节
- 忘记调用 `interfaces.Register`
- 初始化顺序混乱，导致菜单 / 路由 / 表结构缺失

## 真实参考文件

- `server/plugin/announcement/plugin.go`
- `server/plugin/announcement/initialize/router.go`
