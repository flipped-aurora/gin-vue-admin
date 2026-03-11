---
name: model-spec
description: GVA数据模型规范 - 设计模型、处理类型问题时加载
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# GVA 数据模型规范

## 数据模型 (model/xxx.go)

```go
package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Xxx struct {
    global.GVA_MODEL              // ID, CreatedAt, UpdatedAt, DeletedAt
    Name        string  `json:"name" gorm:"column:name;comment:名称"`
    Description string  `json:"description" gorm:"column:description;type:text;comment:描述"`
    Status      int     `json:"status" gorm:"column:status;default:1;comment:状态"`
    Sort        int     `json:"sort" gorm:"column:sort;default:0;comment:排序"`
    CreatedBy   uint    `json:"createdBy" gorm:"column:created_by;comment:创建者ID"`
}

func (Xxx) TableName() string {
    return "xxx_table"
}
```

## 请求模型 (model/request/xxx.go)

```go
package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// CreateXxxRequest 创建请求
type CreateXxxRequest struct {
    Name        string `json:"name" form:"name" binding:"required"`
    Description string `json:"description" form:"description"`
    Status      int    `json:"status" form:"status"`
}

// UpdateXxxRequest 更新请求
type UpdateXxxRequest struct {
    ID          uint   `json:"id" form:"id" binding:"required"`
    Name        string `json:"name" form:"name"`
    Description string `json:"description" form:"description"`
    Status      int    `json:"status" form:"status"`
}

// XxxSearch 搜索请求
type XxxSearch struct {
    request.PageInfo          // 内嵌分页: Page, PageSize
    Name   string `json:"name" form:"name"`
    Status *int   `json:"status" form:"status"` // 指针用于区分0和未传
}
```

## 类型一致性要求

### 同字段类型必须一致

**错误示例**:
```go
// model/xxx.go
type Xxx struct {
    Status int `json:"status"`  // int类型
}

// model/request/xxx.go
type CreateXxxRequest struct {
    Status string `json:"status"`  // 错误！应该也是int
}
```

**正确示例**:
```go
// model/xxx.go
type Xxx struct {
    Status int `json:"status"`
}

// model/request/xxx.go
type CreateXxxRequest struct {
    Status int `json:"status"`  // 保持一致
}
```

## 指针类型处理

### 何时使用指针

| 场景 | 使用指针 | 原因 |
|-----|---------|-----|
| 需要区分零值和未传 | `*int` | 0 vs nil |
| 可选字段 | `*string` | "" vs nil |
| 数据库nullable字段 | `*time.Time` | 支持NULL |

### 转换规则

```go
// 指针 → 非指针 (需要判nil)
func (s *XxxService) ToResponse(m *model.Xxx) response.XxxResponse {
    resp := response.XxxResponse{
        ID: m.ID,
    }
    if m.Name != nil {
        resp.Name = *m.Name
    }
    if m.Status != nil {
        resp.Status = *m.Status
    }
    return resp
}

// 非指针 → 指针 (取地址)
func (s *XxxService) ToModel(req request.CreateXxxRequest) model.Xxx {
    return model.Xxx{
        Name:   &req.Name,
        Status: &req.Status,
    }
}
```

## 常见类型映射

| Go类型 | GORM | JSON | 前端JS |
|-------|------|------|-------|
| `int` | INTEGER | number | number |
| `uint` | UNSIGNED INT | number | number |
| `string` | VARCHAR | string | string |
| `bool` | TINYINT(1) | boolean | boolean |
| `time.Time` | DATETIME | string | string/Date |
| `float64` | DOUBLE | number | number |
| `[]byte` | BLOB | string(base64) | string |

## 字段标签规范

```go
type Example struct {
    // JSON: API序列化名称
    // GORM: 数据库列配置
    // FORM: 表单绑定名称
    // BINDING: 验证规则

    Name string `json:"name" gorm:"column:name;type:varchar(100);not null" form:"name" binding:"required,min=1,max=100"`
}
```

### 常用验证规则

| 规则 | 说明 | 示例 |
|-----|------|-----|
| required | 必填 | `binding:"required"` |
| min/max | 长度/值范围 | `binding:"min=1,max=100"` |
| email | 邮箱格式 | `binding:"email"` |
| oneof | 枚举值 | `binding:"oneof=1 2 3"` |

## 易错点检查清单

- [ ] 同字段在data/request/response模型中类型一致
- [ ] 指针类型有nil检查
- [ ] 数值字段前后端类型匹配 (int vs number)
- [ ] 时间字段格式统一
- [ ] 枚举值定义一致

