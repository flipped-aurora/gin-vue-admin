# API 示例

## 这个文件负责什么

API 层负责接收 HTTP 请求、从合适的位置提取参数、调用 Service、统一返回响应，并补全 Swagger 注释。

## 什么时候应该这样写

- 新增对外接口
- 为某个模块补 CRUD 接口
- 需要统一响应结构和错误处理

## 核心原则

参数从哪里取，不是由“AI 习惯”决定，而是由下面几件事共同决定：

1. 前端怎么传
2. 接口协议怎么约定
3. 当前逻辑到底需要什么数据
4. 这个数据放在哪个位置更合理、更安全

也就是说，API 层不应该机械地一律使用 `ShouldBindJSON`，而应该先判断参数来源，再选择对应取法。

## 常见参数来源与推荐取法

- JSON body: `ShouldBindJSON`
- Query string: `ShouldBindQuery`、`c.Query(...)`、`c.DefaultQuery(...)`
- Path params: `c.Param(...)`
- `multipart/form-data`: `c.FormFile(...)`、`c.DefaultPostForm(...)`、`c.Request.FormValue(...)`
- Header: `c.GetHeader(...)`、`c.Request.Header.Get(...)`
- Cookie: `c.Cookie(...)`

## 推荐写法示例

下面这个示例演示的是 `POST + JSON body` 场景，所以这里使用 `ShouldBindJSON`。

```go
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

// GetOrderList
// @Tags      Order
// @Summary   分页获取订单列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.OrderSearch                                   true  "分页和筛选参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "返回列表、总数、分页信息"
// @Router    /order/getOrderList [post]
func (o *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo systemReq.OrderSearch
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := orderService.GetOrderList(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
```

## 为什么这样写

- API 层统一负责“取参数 + 校验 + 调 Service + 回响应”
- 取参数方式必须与真实数据来源一致，而不是套固定模板
- 例如：
  - 登录、创建、更新这类通常来自 JSON body
  - 列表筛选、分页、导出条件常来自 Query
  - 上传文件通常来自 `multipart/form-data`
  - 鉴权 token、特殊网关头、追踪信息常来自 Header 或 Cookie
- 成功 / 失败统一使用 `response` 包
- Swagger 注释让接口契约对前端和文档生成都可见

## 常见错误

- 在 API 层直接操作数据库
- 直接 `c.JSON(...)`，绕开统一响应
- 没有 Swagger 注释或注释和实际行为不一致
- 不看参数真实来源，机械地一律使用 `ShouldBindJSON`
- 本该从 Header / Cookie / Query / form-data 取的数据，却硬塞进 body

## 真实参考文件

- `server/api/v1/system/sys_user.go`
- `server/api/v1/system/sys_dictionary.go`
- `server/api/v1/system/auto_code_mcp.go`
- `server/api/v1/example/exa_file_upload_download.go`
- `server/utils/claims.go`
