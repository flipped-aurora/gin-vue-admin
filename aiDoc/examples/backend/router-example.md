# Router 示例

## 这个文件负责什么

Router 层负责路由分组、中间件挂载和处理函数绑定，不承载业务逻辑。

## 什么时候应该这样写

- 新增模块路由
- 区分需要操作日志和不需要操作日志的接口
- 对某类路由统一挂载权限或认证中间件

## 推荐写法示例

```go
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (r *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouter := Router.Group("order").Use(middleware.OperationRecord())
	orderRouterWithoutRecord := Router.Group("order")

	{
		orderRouter.POST("createOrder", orderApi.CreateOrder)
		orderRouter.PUT("updateOrder", orderApi.UpdateOrder)
		orderRouter.DELETE("deleteOrder", orderApi.DeleteOrder)
	}
	{
		orderRouterWithoutRecord.POST("getOrderList", orderApi.GetOrderList)
		orderRouterWithoutRecord.GET("findOrder", orderApi.FindOrder)
	}
}
```

## 为什么这样写

- 写操作和读操作分组清晰，方便挂不同中间件
- 路由层只做绑定，职责边界简单明确
- 和项目现有 `InitXxxRouter` 命名方式一致

## 常见错误

- 在路由文件里写业务逻辑
- 所有接口都挂同一种中间件，导致读接口也被记录操作日志
- 直接引用数据库或 Service，而不是引用 API 处理函数

## 真实参考文件

- `server/router/system/sys_user.go`
- `server/router/system/enter.go`
