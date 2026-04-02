# Request 示例

## 这个文件负责什么

请求结构体负责承接前端入参、查询条件和分页参数，不负责数据库关联或业务逻辑。

## 什么时候应该这样写

- 新增列表查询参数
- 新增创建 / 更新接口的请求体
- 需要复用分页、排序、过滤字段时

## 推荐写法示例

```go
package request

import common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type CreateOrderReq struct {
	Name   string `json:"name" form:"name"`
	Status int    `json:"status" form:"status"`
	Remark string `json:"remark" form:"remark"`
}

type OrderSearch struct {
	common.PageInfo
	Name   string `json:"name" form:"name"`
	Status *int   `json:"status" form:"status"`
	OrderKey string `json:"orderKey" form:"orderKey"`
	Desc     bool   `json:"desc" form:"desc"`
}
```

## 为什么这样写

- `json` 与 `form` 标签一起写，方便 JSON 和表单 / 查询参数共用
- 列表查询统一内嵌 `PageInfo`，避免每个模块重复定义分页字段
- 请求结构和数据库模型分离，能减少耦合和误用

## 常见错误

- 在 request 结构里塞数据库关联字段
- 忘记写 `form` 标签，导致查询参数绑定不完整
- 直接复用数据库 model 作为请求体，导致字段过多或语义混乱

## 真实参考文件

- `server/model/system/request/sys_user.go`
- `server/model/common/request/common.go`
