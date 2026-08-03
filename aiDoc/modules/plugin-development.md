# 插件开发约束

## 适用范围与编辑前门禁

本规则适用于以下目录中的新增、修改、重构和生成工作：

- `server/plugin/<name>/`
- `web/src/plugin/<name>/`
- `server/resource/plugin/server/`
- `server/resource/plugin/web/`

开始任何文件编辑前，AI 必须完成以下动作：

1. 阅读本文件与 `aiDoc/examples/plugin/full-plugin-example.md`
2. 后端插件任务阅读 `aiDoc/examples/backend/plugin-go-example.md`
3. 根据涉及层阅读 `aiDoc/modules/backend-layer-rules.md`、`aiDoc/frontend-backend/frontend-rules.md`、`aiDoc/frontend-backend/boundary.md` 中的相关规则
4. 阅读涉及层在 `server/resource/plugin/server/` 或 `server/resource/plugin/web/` 下的当前生成模板
5. 阅读对应的后端或前端分层示例；修改现有插件时还要阅读目标插件的相关文件，但不能把目标插件自动视为规范来源
6. 完成上述必读清单后，在首个且位于文件编辑前的工作更新中列出本次已读取的规则、模板和示例；清单未完成前不得编辑插件代码

## 参考优先级与边界

插件任务的参考优先级固定如下：

1. `AGENTS.md`
2. 本文件及相关模块、分层、前后端边界规则
3. 当前 `server/resource/plugin/` 生成模板
4. `aiDoc/examples/` 中对应的讲解型示例
5. 示例文档明确列出的真实参考文件

使用真实插件时必须遵守以下边界：

- 不得随机选择“看起来相近”或“当前能运行”的插件整目录复制
- `server/plugin/email/` 是 v1 遗留实现，不得作为新插件或 v2 插件参考
- `server/plugin/plugin-tool/` 是内部工具，不得作为普通业务插件参考
- `announcement` 只用于 `aiDoc/examples/plugin/full-plugin-example.md` 明确列出的入口、路由和 `enter.go` 聚合职责；不得把其未列出的历史业务层写法扩展为规范
- 真实代码、模板、示例互相冲突时，必须按上述优先级判断；高优先级规则已经明确时直接按其实现，并同步修正过时示例，不得静默沿用低优先级旧写法

## 后端插件结构

标准后端插件必须按职责保持以下结构；确需偏离时，必须在编辑前说明原因并取得用户确认：

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

标准前端插件必须按职责保持以下结构；确需偏离时，必须在编辑前说明原因并取得用户确认：

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

- 保持自包含
- 保持可配置
- 预留扩展点
- 与主系统保持一致的风格与约定

## 标准开发流程

1. 先明确插件边界与数据模型
2. 先完成后端模型、服务、接口与初始化
3. 再完成前端接口封装、页面与表单
4. 最后完成菜单、权限、联调与测试

## 交付检查表

插件代码交付前必须逐项核对：

- 后端、前端目录职责与 `enter.go` 分组聚合符合标准结构，没有把跨层逻辑堆进单文件
- 新插件使用 v2 `interfaces.Plugin`，在 `init()` 中自注册，根 `plugin.go` 只负责注册与初始化调度
- 私有路由中间件完整且顺序为 `JWTAuth -> MustChangePwdGuard -> CasbinHandler -> DataScope`
- API 层把 `c.Request.Context()` 传入 Service，所有相关 GORM 操作使用 `WithContext(ctx)`
- 分页请求使用 `request.PageInfo` 与 `LimitOffset()`，Swagger 列表响应声明具体元素类型
- 前后端请求字段、响应结构与路由一致，前端保持 `api/`、`view/`、`form/` 职责分离并遵循 UnoCSS 规则
- 通用插件规则或标准结构发生变化时，同步更新 `server/resource/plugin/` 对应模板，避免后续继续生成旧代码
- 已运行与改动范围匹配的格式化、单元测试、构建或页面验证，并检查完整 diff
- 相关跨栈契约、插件规则、示例和业务记忆已同步更新
