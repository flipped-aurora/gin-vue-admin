# Service 示例

## 这个文件负责什么

Service 层负责业务逻辑、数据库查询、事务控制和数据拼装，不负责 HTTP 参数绑定和响应输出。

## 什么时候应该这样写

- 新增列表查询、创建、更新、删除逻辑
- 需要事务操作
- 需要对 model 和 request 做转换

## 推荐写法示例

```go
package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type OrderService struct{}

func (s *OrderService) GetOrderList(ctx context.Context, info systemReq.OrderSearch) (list []system.Order, total int64, err error) {
	limit, offset := info.LimitOffset()

	db := global.GVA_DB.WithContext(ctx).Model(&system.Order{})
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error
	return list, total, err
}
```

## 当表带数据权限时怎么写

数据权限（行级过滤）由统一引擎的 GORM 全局回调实现（`server/utils/datascope/`），
Service 层**不写任何范围过滤条件**。以下示例均摘自真实代码，直接照此写。

查询——和普通表完全一样（`server/service/example/exa_customer.go` 的 `GetCustomerInfoList`）：

```go
func (exa *CustomerService) GetCustomerInfoList(ctx context.Context, sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	_ = sysUserAuthorityID // 数据范围已由数据权限引擎接管, 旧的按角色 DataAuthorityId 手写过滤已移除
	limit, offset := info.LimitOffset()
	// 数据范围(本部门及子级/仅本人)由数据权限引擎的 GORM 回调按当前角色 data_scope
	// 自动追加 dept_id/created_by 过滤, 此处只写常规分页查询即可。
	db := global.GVA_DB.WithContext(ctx).Model(&example.ExaCustomer{})
	var CustomerList []example.ExaCustomer
	if err = db.Count(&total).Error; err != nil {
		return CustomerList, total, err
	}
	err = db.Limit(limit).Offset(offset).Preload("SysUser").Find(&CustomerList).Error
	return CustomerList, total, err
}
```

创建同样零特殊代码（`CreateExaCustomer` 就是一句普通 `Create`），`created_by`/`dept_id`
由 create 回调自动盖章；更新时用 `Omit` 保护归属列（`UpdateExaCustomer`）：

```go
func (exa *CustomerService) UpdateExaCustomer(ctx context.Context, e *example.ExaCustomer) (err error) {
	// 归属列(dept_id/created_by)创建时盖章后即不可变, 更新时忽略, 防止被表单未回传的零值覆盖
	err = global.GVA_DB.WithContext(ctx).Omit("dept_id", "created_by").Save(e).Error
	return err
}
```

两种合法旁路（均摘自真实代码）：

```go
// 1) 显式旁路：引擎自身或确需跨范围的查询
//    （server/service/system/data_scope.go 的 BuildIdentity，旁路以避免回调递归）
err := global.GVA_DB.WithContext(ctx).Set("data_scope:skip", true).
	Select("authority_id", "data_scope").
	Where("authority_id = ?", authorityID).First(&auth).Error

// 2) 系统上下文：定时任务/CLI/初始化等无请求身份的场景（server/initialize/timer.go）
sysDB := global.GVA_DB.WithContext(datascope.WithSystem(context.Background()))
```

必须遵守的规则：

- **`WithContext(ctx)` 不可省，且 ctx 必须一路来自 `c.Request.Context()`**：身份由 DataScope 中间件注入 request context，引擎回调从中读取；漏传 ctx 等于旁路数据权限（现阶段会放行，但响亮告警并落审计表 `sys_data_access_logs`）
- **不要手写 `dept_id` / `created_by` 的权限过滤 WHERE**：引擎会按角色档位自动追加，手写会重复甚至冲突
- **创建时不要手动赋值 `CreatedBy` / `DeptId`**：引擎在 create 回调里按当前身份自动盖章
- **`UpdatedBy` / `DeletedBy` 同样不手动赋值**：表有 `updated_by` 列则每次更新自动盖；表同时有 `deleted_by` 列与 `gorm.DeletedAt` 时，软删除的那条 UPDATE 会同时 SET `deleted_at` + `deleted_by`（硬删除 / `Unscoped` 不盖）。盖章只在有请求身份时发生：无身份、`WithSystem` 系统上下文、`UpdateColumn`（SkipHooks）等场景静默不盖
- **更新走 `Save` 等全量写时 `Omit("dept_id", "created_by")`**：归属列创建盖章后不可变，防止被表单未回传的零值覆盖；`updated_by` 不要放进 Omit，它靠引擎每次更新自动刷新
- 确需跨数据范围查询（如部门树自身的计算），用 `db.Set("data_scope:skip", true)` 显式旁路
- 定时任务 / CLI / 初始化等无请求身份的系统场景，用 `datascope.WithSystem(ctx)` 标记系统上下文放行，不要裸用 `context.Background()`

## 为什么这样写

- 只返回业务结果和 `error`，便于 API 层统一处理
- 以 `ctx context.Context` 为首参并用 `global.GVA_DB.WithContext(ctx)` 查询，串联请求链路（API 层传 `c.Request.Context()`）
- 分页统一用 `info.LimitOffset()`（`request.PageInfo` 提供），pageSize 超过 `MaxPageSize`(100) 会自动截断，不要手写换算
- 直接使用 `global.GVA_DB` 或事务对象，保持项目一致性
- 查询条件尽量在 Service 层收敛，不散落到 API 或 Router

## 常见错误

- 在 Service 层接触 `gin.Context`
- 手写 `limit := info.PageSize; offset := info.PageSize * (info.Page - 1)` 换算，绕过 `LimitOffset()` 的上限保护
- 给数据权限表手写 `dept_id`/`created_by` 过滤条件，或漏传 ctx 导致引擎拿不到身份（旁路数据权限）
- 在 Service 层直接拼 `response.OkWith...`
- 把参数绑定、权限判断、响应格式化和业务逻辑混在一起
- 涉及多表更新却不使用事务

## 真实参考文件

- `server/service/system/sys_user.go`
- `server/service/example/exa_customer.go`（`GetCustomerInfoList` 是数据权限表的标准查询范例）
- `server/utils/datascope/datascope.go`（数据权限引擎：档位、护栏与旁路机制）
