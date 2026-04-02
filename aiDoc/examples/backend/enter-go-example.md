# enter.go 示例

## 这个文件负责什么

`enter.go` 用来聚合某一层的模块入口，降低循环引用风险，让其他层通过统一入口访问能力。

## 什么时候应该这样写

- 新增一个模块组
- 需要把多个 Router / API / Service 聚合到统一入口
- 需要让其他层通过 `GroupApp` 访问当前层能力

## 推荐写法示例

### 根级聚合

```go
package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
```

### 模块级聚合

```go
package system

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	UserRouter
	OrderRouter
}

var (
	baseApi  = api.ApiGroupApp.SystemApiGroup.BaseApi
	orderApi = api.ApiGroupApp.SystemApiGroup.OrderApi
)
```

## 为什么这样写

- 统一入口比跨文件直接互相引用更稳定
- 可以把常用 API / Service 别名集中在这里
- 模块扩展时，只需在 `enter.go` 增加聚合字段

## 常见错误

- 不用 `enter.go`，跨层直接互相 import
- 在多个文件里重复声明同一组别名变量
- 新增模块后忘记把它注册进 `Group` 结构体

## 真实参考文件

- `server/api/v1/enter.go`
- `server/service/enter.go`
- `server/router/system/enter.go`
