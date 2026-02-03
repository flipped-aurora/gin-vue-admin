---
name: backend-dev
description: GVA后端开发规范 - 开发Go后端、API接口时加载
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# GVA 后端开发规范

## 核心原则

### 1. 严格的分层架构

**依赖链**: `Router → API → Service → Model`

- **Router层**: 定义路由规则，映射到API处理函数
- **API层**: HTTP请求入口，参数校验，调用Service，返回响应
- **Service层**: 核心业务逻辑，数据库CRUD操作
- **Model层**: 数据库表结构定义，请求/响应DTO

**禁止**: 跨层调用。API层不能直接操作数据库，Service层不能处理`gin.Context`。

### 2. enter.go 组管理模式

所有 api/service/router 层**必须**使用 `enter.go` 管理：

```go
// service/enter.go
package service

type ServiceGroup struct {
    XxxService
    YyyService
}

var ServiceGroupApp = new(ServiceGroup)
```

```go
// api/enter.go
package api

type ApiGroup struct {
    XxxApi
    YyyApi
}

var ApiGroupApp = new(ApiGroup)
```

```go
// router/enter.go
package router

type RouterGroup struct {
    XxxRouter
    YyyRouter
}

var RouterGroupApp = new(RouterGroup)
```

### 3. 模块间引用关系

```go
// API层引用Service层
var xxxService = service.ServiceGroupApp.XxxService

// Router层引用API层
api.ApiGroupApp.XxxApi.XxxMethod

// Initialize引用Router层
router.RouterGroupApp.XxxRouter.InitXxxRouter
```

## 各层实现规范

### Model层 (`model/`)

```go
// model/xxx.go - 数据模型
type Xxx struct {
    global.GVA_MODEL  // 包含 ID, CreatedAt, UpdatedAt
    Name        string  `json:"name" gorm:"column:name;comment:名称"`
    Status      int     `json:"status" gorm:"column:status;default:1"`
}

func (Xxx) TableName() string {
    return "xxx_table"
}
```

```go
// model/request/xxx.go - 请求模型
type CreateXxxRequest struct {
    Name   string `json:"name" form:"name" binding:"required"`
    Status int    `json:"status" form:"status"`
}

type XxxSearch struct {
    request.PageInfo        // 内嵌分页
    Name string `json:"name" form:"name"`
}
```

### Service层 (`service/`)

```go
// service/xxx_service.go
type XxxService struct{}

// CreateXxx 创建记录
func (s *XxxService) CreateXxx(data model.Xxx) error {
    return global.GVA_DB.Create(&data).Error
}

// GetXxxList 分页查询
func (s *XxxService) GetXxxList(info request.XxxSearch) (list []model.Xxx, total int64, err error) {
    db := global.GVA_DB.Model(&model.Xxx{})
    if info.Name != "" {
        db = db.Where("name LIKE ?", "%"+info.Name+"%")
    }
    err = db.Count(&total).Error
    if err != nil {
        return
    }
    err = db.Limit(info.PageSize).Offset(info.PageSize * (info.Page - 1)).Find(&list).Error
    return
}
```

### API层 (`api/`)

```go
// api/xxx_api.go
type XxxApi struct{}

var xxxService = service.ServiceGroupApp.XxxService

func (a *XxxApi) CreateXxx(c *gin.Context) {
    var req request.CreateXxxRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }

    data := model.Xxx{Name: req.Name, Status: req.Status}
    if err := xxxService.CreateXxx(data); err != nil {
        response.FailWithMessage("创建失败: "+err.Error(), c)
        return
    }
    response.OkWithMessage("创建成功", c)
}
```

### Router层 (`router/`)

```go
// router/xxx_router.go
type XxxRouter struct{}

func (r *XxxRouter) InitXxxRouter(Router *gin.RouterGroup) {
    xxxRouter := Router.Group("xxx").Use(middleware.OperationRecord())
    xxxRouterWithoutRecord := Router.Group("xxx")

    api := v1.ApiGroupApp.XxxApi
    {
        xxxRouter.POST("create", api.CreateXxx)
        xxxRouter.DELETE("delete", api.DeleteXxx)
        xxxRouter.PUT("update", api.UpdateXxx)
    }
    {
        xxxRouterWithoutRecord.GET("list", api.GetXxxList)
        xxxRouterWithoutRecord.GET("find", api.FindXxx)
    }
}
```

## 统一响应格式

```go
// 成功响应
response.Ok(c)
response.OkWithMessage("操作成功", c)
response.OkWithData(data, c)
response.OkWithDetailed(data, "获取成功", c)

// 失败响应
response.Fail(c)
response.FailWithMessage("操作失败", c)
response.FailWithDetailed(nil, "详细错误信息", c)
```

## 参考示例

经典插件结构: `server/plugin/announcement/`

