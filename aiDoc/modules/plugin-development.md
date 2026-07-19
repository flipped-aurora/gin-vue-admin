# 插件开发约束

## 后端插件结构

后端插件推荐保持以下结构：

- `api/`
- `config/`
- `gen/`（gorm/gen 代码生成入口，`go:generate` 独立程序）
- `initialize/`
- `model/`
- `model/request/`
- `plugin/`（插件内全局配置访问包：`var Config config.Config`）
- `router/`
- `service/`
- `plugin.go`

## 前端插件结构

前端插件推荐保持以下结构：

- `api/`
- `form/`
- `view/`

以上与前端插件生成模板（`server/resource/plugin/web/`）的产物一致；确有需要时可自建 `components/` 等子目录，但不是模板产物、不做强制要求。

## 插件入口约束

`plugin.go` 至少要承担以下职责：

- 实现 v2 插件接口 `interfaces.Plugin`（`server/utils/plugin/v2`，只有一个 `Register(group *gin.Engine)` 方法）
- 在 `init()` 中调用 `interfaces.Register(Plugin)` 完成自注册
- 在 `Register` 中完成路由挂载，并按需调度 `initialize` 包的 `Gorm / Api / Menu / Dictionary / Viper` 初始化

> 遗留的 v1 接口（带 `RouterPath()` 方法，仅 email 插件仍在使用，由 `plugin_biz_v1.go` 手动挂载）不要用于新插件。

## 路由注册约束

v2 插件在各自 `initialize/router.go` 中从 `*gin.Engine` 自建 public/private 组。私有组的中间件链必须与主系统 PrivateGroup（`server/initialize/router.go`）完全对齐、顺序一致：

```go
private.Use(middleware.JWTAuth()).Use(middleware.MustChangePwdGuard()).Use(middleware.CasbinHandler()).Use(middleware.DataScope())
```

- `MustChangePwdGuard`：强制改密守卫。JWT 携带 `MustChangePwd=true` 时仅放行改密/用户信息/登出接口，其余一律 403
- `DataScope`：行级数据权限身份注入。依据 claims 构建数据权限身份并写入 `c.Request.Context()`，供 Service 层 `WithContext(ctx)` 透传到 GORM 全局回调消费
- 缺了这两个中间件：插件接口会绕过强制改密拦截，且数据权限身份不会注入，行级数据过滤对插件接口失效
- v1 遗留插件（email）挂在主 PrivateGroup 上，自动继承完整中间件链，无需在插件内重复挂载
- 插件代码生成模板 `server/resource/plugin/server/initialize/router.go.tpl` 同步保持该链，不要回退成旧的两件套

## 插件设计原则

- 尽量自包含
- 保持可配置
- 预留扩展点
- 与主系统保持一致的风格与约定

## 推荐开发流程

1. 先明确插件边界与数据模型
2. 先完成后端模型、服务、接口与初始化
3. 再完成前端接口封装、页面与表单
4. 最后完成菜单、权限、联调与测试
