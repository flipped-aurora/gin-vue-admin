---
name: plugin-dev
description: GVA插件开发规范 - 开发新插件时加载
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# GVA 插件开发规范

## 后端插件结构

```
server/plugin/[插件名]/
├── api/                # API控制器
│   ├── enter.go       # API组入口
│   └── xxx.go         # 具体API
├── config/            # 插件配置
│   └── config.go
├── initialize/        # 初始化模块
│   ├── api.go        # API注册
│   ├── gorm.go       # 数据库迁移
│   ├── menu.go       # 菜单初始化
│   ├── router.go     # 路由初始化
│   └── viper.go      # 配置加载
├── model/             # 数据模型
│   ├── xxx.go        # 数据库模型
│   └── request/      # 请求模型
├── router/            # 路由定义
│   ├── enter.go      # 路由组入口
│   └── xxx.go        # 具体路由
├── service/           # 业务服务
│   ├── enter.go      # 服务组入口
│   └── xxx.go        # 具体服务
└── plugin.go          # 插件入口
```

## 前端插件结构

```
web/src/plugin/[插件名]/
├── api/               # API接口
│   └── xxx.js
├── components/        # 插件组件
│   └── XxxComp.vue
├── view/             # 插件页面
│   └── index.vue
├── form/             # 表单组件
│   └── XxxForm.vue
└── config.js         # 插件配置
```

## 核心文件实现

### plugin.go (插件入口)

```go
package xxx

import (
    "github.com/gin-gonic/gin"
    "github.com/xxx/server/plugin/xxx/initialize"
)

type XxxPlugin struct{}

func CreateXxxPlugin() *XxxPlugin {
    return &XxxPlugin{}
}

// Register 实现插件接口
func (p *XxxPlugin) Register(group *gin.RouterGroup) {
    initialize.InitializeRouter(group)
}

// RouterPath 返回插件路由根路径
func (p *XxxPlugin) RouterPath() string {
    return "xxx"
}
```

### initialize/gorm.go (数据库迁移)

```go
package initialize

import (
    "github.com/xxx/server/global"
    "github.com/xxx/server/plugin/xxx/model"
)

func InitializeDB() error {
    return global.GVA_DB.AutoMigrate(
        &model.Xxx{},
        &model.XxxDetail{},
    )
}
```

### initialize/router.go (路由注册)

```go
package initialize

import (
    "github.com/gin-gonic/gin"
    "github.com/xxx/server/plugin/xxx/router"
)

func InitializeRouter(group *gin.RouterGroup) {
    routerGroup := router.RouterGroupApp
    routerGroup.XxxRouter.InitXxxRouter(group)
}
```

### initialize/menu.go (菜单初始化)

```go
package initialize

import (
    "github.com/xxx/server/global"
    "github.com/xxx/server/model/system"
)

func InitializeMenu() error {
    // 创建父菜单
    parentMenu := system.SysBaseMenu{
        Path:      "xxx",
        Name:      "xxx",
        Component: "view/xxx/index.vue",
        Meta: system.Meta{
            Title: "XXX管理",
            Icon:  "document",
        },
    }

    if err := global.GVA_DB.Create(&parentMenu).Error; err != nil {
        return err
    }

    // 创建子菜单...
    return nil
}
```

## 开发工作流

### 1. 创建插件目录结构

```bash
mkdir -p server/plugin/xxx/{api,config,initialize,model/request,router,service}
mkdir -p web/src/plugin/xxx/{api,view,components,form}
```

### 2. 后端开发顺序

1. **Model**: 定义数据模型和请求模型
2. **Service**: 实现业务逻辑
3. **API**: 创建控制器
4. **Router**: 配置路由
5. **Initialize**: 编写初始化代码
6. **plugin.go**: 创建插件入口

### 3. 前端开发顺序

1. **API**: 封装后端接口
2. **View**: 创建页面组件
3. **Components**: 抽取可复用组件
4. **Form**: 创建表单组件

### 4. 注册插件

在 `server/initialize/plugin.go` 中注册：

```go
import xxxPlugin "github.com/xxx/server/plugin/xxx"

func InstallPlugin(Router *gin.Engine) {
    // ...
    PluginInit(Router, xxxPlugin.CreateXxxPlugin())
}
```

## 质量标准

- [ ] 功能完整：满足业务需求
- [ ] 代码规范：遵循GVA架构
- [ ] 类型一致：前后端字段类型匹配
- [ ] 文档完整：Swagger注释齐全
- [ ] 权限控制：菜单和API权限配置

## 参考示例

官方示例插件: `server/plugin/announcement/`

