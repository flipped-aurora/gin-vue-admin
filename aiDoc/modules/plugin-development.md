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
