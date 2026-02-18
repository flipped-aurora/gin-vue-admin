---
name: swagger-spec
description: GVA Swagger文档规范 - 编写API文档时加载
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
---
name: swagger-spec
description: GVA Swagger文档规范 - 编写API文档时加载
---

# GVA Swagger 文档规范

## 注释格式

```go
// CreateXxx 创建XXX
// @Tags      XxxModule
// @Summary   创建一个新的XXX
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body request.CreateXxxRequest true "创建参数"
// @Success   200 {object} response.Response{msg=string} "创建成功"
// @Failure   400 {object} response.Response "请求参数错误"
// @Router    /xxx/create [post]
func (a *XxxApi) CreateXxx(c *gin.Context) {
    // ...
}
```

## 标签说明

| 标签 | 说明 | 示例 |
|-----|------|-----|
| @Tags | 接口分组 | `@Tags UserModule` |
| @Summary | 接口简述 | `@Summary 获取用户列表` |
| @Description | 详细描述 | `@Description 分页获取用户列表，支持按名称搜索` |
| @Security | 认证方式 | `@Security ApiKeyAuth` |
| @accept | 请求格式 | `@accept application/json` |
| @Produce | 响应格式 | `@Produce application/json` |
| @Param | 请求参数 | 见下方详解 |
| @Success | 成功响应 | `@Success 200 {object} response.Response` |
| @Failure | 失败响应 | `@Failure 400 {object} response.Response` |
| @Router | 路由路径 | `@Router /user/list [get]` |

## @Param 参数详解

格式: `@Param 参数名 位置 类型 必填 "描述"`

### 位置类型

| 位置 | 说明 | 示例 |
|-----|------|-----|
| body | 请求体JSON | `@Param data body request.Xxx true "请求体"` |
| query | URL查询参数 | `@Param page query int false "页码"` |
| path | URL路径参数 | `@Param id path int true "记录ID"` |
| header | 请求头 | `@Param Authorization header string true "Token"` |
| formData | 表单数据 | `@Param file formData file true "上传文件"` |

## 完整示例

### 创建接口

```go
// CreateXxx 创建XXX
// @Tags      Xxx管理
// @Summary   创建XXX记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body request.CreateXxxRequest true "名称和描述"
// @Success   200 {object} response.Response{msg=string} "创建成功"
// @Router    /xxx/create [post]
func (a *XxxApi) CreateXxx(c *gin.Context) {}
```

### 删除接口

```go
// DeleteXxx 删除XXX
// @Tags      Xxx管理
// @Summary   删除XXX记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body request.IdsReq true "ID列表"
// @Success   200 {object} response.Response{msg=string} "删除成功"
// @Router    /xxx/delete [delete]
func (a *XxxApi) DeleteXxx(c *gin.Context) {}
```

### 更新接口

```go
// UpdateXxx 更新XXX
// @Tags      Xxx管理
// @Summary   更新XXX记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body request.UpdateXxxRequest true "更新内容"
// @Success   200 {object} response.Response{msg=string} "更新成功"
// @Router    /xxx/update [put]
func (a *XxxApi) UpdateXxx(c *gin.Context) {}
```

### 查询单条

```go
// FindXxx 查询XXX
// @Tags      Xxx管理
// @Summary   根据ID查询XXX
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id query int true "记录ID"
// @Success   200 {object} response.Response{data=model.Xxx} "查询成功"
// @Router    /xxx/find [get]
func (a *XxxApi) FindXxx(c *gin.Context) {}
```

### 分页列表

```go
// GetXxxList 获取XXX列表
// @Tags      Xxx管理
// @Summary   分页获取XXX列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body request.XxxSearch true "分页和搜索参数"
// @Success   200 {object} response.Response{data=response.PageResult{list=[]model.Xxx}} "获取成功"
// @Router    /xxx/list [post]
func (a *XxxApi) GetXxxList(c *gin.Context) {}
```

## 生成文档

```bash
cd server
swag init
```

访问文档: `http://localhost:8888/swagger/index.html`

## 注意事项

1. 每个API函数**必须**有完整注释
2. @Tags 使用中文便于阅读
3. @Param 的类型要与实际结构体匹配
4. @Success 的 data 类型要准确反映返回值

