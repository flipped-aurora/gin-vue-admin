# AGENT.MD

## 目的

本文件是本仓库内 AI 协作规则的唯一真源。

`.codex/`、`.claude/`、`.cursor/`、`.trae/` 下的规则文件仅作为兼容适配层，不能再次演变成各自独立维护的 project rule 副本。

## 读取顺序

按下面顺序加载项目上下文：

1. `AGENT.MD`
2. `aiDoc/README.md`
3. 按任务读取以下目录中的相关文件：
   - `aiDoc/relations/`
   - `aiDoc/modules/`
   - `aiDoc/frontend-backend/`
   - `aiDoc/examples/`
   - `aiDoc/memory/`
4. 仅在当前工具确实依赖时，再读取工具目录下的适配文件

若内容冲突，以 `AGENT.MD` 为准。

## 仓库概览

- `server/`: Go + Gin 后端
- `web/`: Vue 3 + Vite 前端
- `deploy/`: 部署资产
- `docs/`: 项目文档与设计记录
- `aiDoc/`: AI 协作文档层

## 工程规则

### 架构

- 保持现有后端分层：`Router -> API -> Service -> Model`
- `enter.go` 继续作为分组注册与组合入口
- API 层处理 HTTP 相关逻辑，Service 层不要依赖 `gin.Context`
- 对外接口的 Swagger 注释必须和真实行为保持一致
- Swagger 的 `@Success` 响应要落到具体类型，让 swag 能生成真实返回结构：列表用 `response.PageResult{list=[]Model}`、详情用具体 model，不要停留在空的 `response.PageResult` 或 `data=object`（仅动态/示例数据才用 `object`）；细则见 `aiDoc/modules/backend-layer-rules.md`
- 列表分页统一走 `request.PageInfo`：Service 层取 limit/offset 一律用 `info.LimitOffset()`（内置 `MaxPageSize=100` 截断），不要手写 `PageSize*(Page-1)` 换算
- `CreatedBy/UpdatedBy/DeletedBy/DeptId` 公共操作字段（`dept_id` 为数据权限的归属部门列）仅在业务表需要数据权限时才创建（对应代码生成器的 AutoCreateResource），手写时不要自造 `CreatorID` 之类同义字段；细则见 `aiDoc/modules/backend-layer-rules.md`
- 行级数据权限由统一引擎的 GORM 全局回调自动过滤与盖章，Service 只负责把 `c.Request.Context()` 一路透传（`WithContext(ctx)`），不手写 `dept_id`/`created_by` 范围条件；细则见 `aiDoc/examples/backend/service-example.md`
- 优先沿用 gin-vue-admin 现有模式，不做无关的大改

### 前后端协作

- 明确请求与响应契约
- 保持统一响应结构：`{ code, data, msg }`
- 保持统一分页结构：`{ page, pageSize, total, list }`
- 前后端字段名和数据类型保持一致
- 优先复用 `web/src/utils/` 里的工具函数
- 涉及跨栈边界变更时，同步更新 `aiDoc/frontend-backend/`

### 前端样式

- 前端已接入 UnoCSS（`web/uno.config.js`，presetWind3）。**写样式优先用 UnoCSS 原子类**：布局、间距、排版、尺寸、颜色等能用原子类表达的，一律用原子类，不再新增自定义 class
- 只有以下情况才写 `<style scoped>`：需要 `:deep()` 覆盖 Element Plus 内部样式、伪类/伪元素、复杂选择器，或原子类确实难以表达的样式
- 避免内联 `style`；动态样式用绑定的原子 class 或 CSS 变量，主题相关能力优先走 CSS 变量
- 细则见 `aiDoc/frontend-backend/frontend-rules.md`

### 图标

- **菜单图标**(`server/source/system/menu.go` 的 `Icon` 字段)优先用空心(线框)风格,避免填充式/实心款(如 Element Plus 的 `*-filled`,及 `avatar`/`platform`/`management`/`lock` 等视觉实心款);**找不到合适的空心图标时,去 Iconify(优先 `lucide`)挑一个合适的空心 svg、规整后存为本地 `web/src/assets/icons/<name>-gva.svg`(不要手画自己发挥,也不要将就用实心款)**;由 `web/src/core/global.js` 自动注册、菜单按名引用,新增后需重启/重新构建前端以重生成 sprite(线宽统一由 `svgIcon.vue` 提供,**根 svg 勿写 `stroke-width`**——构建插件会破坏它)
- **其它系统 / 业务开发的图标**:只要语义合适即可,**不必在意空心还是实心**;优先复用现有图标集(lucide 等),确实没有合适的再自建 svg
- 自定义 SVG 的规格、机制与现有图标清单见 `aiDoc/frontend-backend/frontend-rules.md` 的「图标规范」

### 页面点触测试(AI 驱动浏览器验证)

- 前端页面改动需要真实浏览器点触验证时,登录态获取按 `aiDoc/frontend-backend/page-click-testing.md` 执行:优先静默读取 `.local/gva-test-token`(已 gitignore);没有或失效时按该文档的标准话术向用户索取 token,注入 localStorage 后点触
- 点触依赖浏览器自动化能力:当前环境不具备时,**主动建议用户安装**(按该文档「环境前置」一节的话术给出推荐项,首选 Playwright),经用户确认后再装,不要不问就装;用户拒绝则回退为"给出人工目测清单"
- 需要覆盖登录链路本身(验证码/锁定)时,用「系统设置 → 安全配置」的验证码阈值临时直登,测完改回
- token 是真实凭证:不写入任何会提交的文件、不出现在截图/日志/commit 信息里;点触造成的破坏性数据操作先征得用户同意

### 后端测试

- 后端单元测试需要 gva 全局单例(`global.GVA_DB`/`GVA_LOG`/`GVA_CACHE`/`GVA_REDIS`/`GVA_CONFIG`)时，统一复用 `server/internal/testutil`，不要在各 `_test.go` 里重复手写 sqlite `:memory:`、`AutoMigrate`、nop logger、存旧值/赋值/`t.Cleanup` 还原这套样板
- 常用入口：`testutil.NewMemoryDB(t, models...)`(内存库 + AutoMigrate + 赋值 `GVA_DB` + cleanup 还原，可选 `WithDataScopeCallbacks()`/`WithoutAutoMigrate()`)、`InitMemoryCache(t, 0)`、`InitNopLogger()`、`NewRedisOrSkip(t)`(无 Redis 自动 skip)、`LoadConfig(t, path)` / `LoadDefaultTestConfig(t)`；每个 helper 的语义与选项以包内 godoc 为准
- 约定：helper 默认把新建实例赋值到对应全局单例并在 `t.Cleanup` 还原(个别如 Redis 需显式开启)；`t.Parallel()` 并行测试改用 `*WithoutGlobal` 变体或直接持有局部实例，避免全局单例数据竞争
- `testutil` 仅供测试使用，置于 `server/internal/` 下，只有 server 模块内可 import；新增测试优先向 `testutil` 补 helper，不要另起一套内联样板
- 参考用法：`server/service/system/sys_timed_task_runner_test.go`

### 插件与模块

- 后端插件放在 `server/plugin/<name>/`
- 前端插件放在 `web/src/plugin/<name>/`
- 插件私有路由组的中间件链必须与主系统 PrivateGroup（`server/initialize/router.go`）对齐且顺序一致：`JWTAuth -> MustChangePwdGuard -> CasbinHandler -> DataScope`；v2 插件在各自 `initialize/router.go` 显式挂载，插件代码生成模板同样适用；细则见 `aiDoc/modules/plugin-development.md`
- 稳定的模块职责、入口和边界说明放到 `aiDoc/modules/`

### 示例文档

- `aiDoc/examples/` 是讲解型示例层
- 示例文档不是要求逐字复制，而是告诉 AI 该如何按项目标准组织代码
- 当 AI 需要新增某一层文件时，应先阅读对应示例，再开始实现

### 记忆规则

- `aiDoc/memory/long-term/` 用于记录长期稳定的用户偏好、协作方式和跨任务约束
- `aiDoc/memory/business/` 用于记录每次用户提出的业务需求
- 这里的“业务需求”指新增或修改模块、接口、页面、流程、业务规则、插件能力等项目需求
- **一条记忆的粒度是「一个功能点 / 一次独立的变更意图」，不是「一个大模块」**：同一大模块下的不同功能点必须各自独立成一个文件，禁止用“新增子需求”的方式把新功能追加进已有文件
- 用户提出新的功能点时，AI 必须新建一个 `business` 记忆文件，并在 `demand-index.md` 中登记
- 只有更新**同一条记忆**（同一个功能点）的状态、进展或已确认约束时，才编辑那个已有文件
- 同一大模块下的多个功能文件，用统一文件名前缀归组（如 `cli-xxx.md`），并在文件之间用反向链接关联上下文
- 当某条业务需求沉淀成长期稳定偏好时，再提炼到 `long-term`
- 长期记忆按主题一条一文件维护，业务记忆按功能点一条一文件维护
- `README.md`、`project-memory.md`、`demand-index.md` 这类文件只承担说明或索引职责，不承载多条记忆正文
- 更新记忆时，只修改对应那条记忆文件与必要索引，不要通过重写汇总正文来覆盖其他历史记忆

### 文档维护

- `AGENT.MD` 只保留高层、稳定、跨任务的规则
- 细节上下文写入 `aiDoc/`
- 结构关系放在 `aiDoc/relations/`
- 示例写法放在 `aiDoc/examples/`
- 长期记忆与业务记忆放在 `aiDoc/memory/`
- 若项目级 AI 规则发生变化，先改 `AGENT.MD`，再按需更新适配层

### 版权与授权保护规则

- 涉及版权声明、作者署名、许可证、商用授权、品牌标识及其相关实现的任务，必须遵循仓库内公开的授权与发布规则
- 具体判定、处理边界和协作方式，统一参考 `aiDoc/relations/branding.md`
- 若公开规则与口头说明冲突，以仓库内可审计的文档、配置和代码事实为准

### 代码读取约束

- 无论什么情况，都不要直接读取 `node_modules/` 中的代码
- 如需了解第三方库行为，优先查看项目源码中的调用方式、锁文件、配置文件、官方文档或包的公开类型/说明文件，而不是进入 `node_modules/` 逐文件读取实现

## AI 文档索引

- `aiDoc/README.md`
- `aiDoc/relations/repo-profile.md`
- `aiDoc/relations/development-workflow.md`
- `aiDoc/relations/system-map.md`
- `aiDoc/relations/branding.md`
- `aiDoc/modules/module-index.md`
- `aiDoc/modules/backend-layer-rules.md`
- `aiDoc/modules/plugin-development.md`
- `aiDoc/frontend-backend/boundary.md`
- `aiDoc/frontend-backend/frontend-rules.md`
- `aiDoc/frontend-backend/page-click-testing.md`
- `aiDoc/frontend-backend/frontend-utils.md`
- `aiDoc/frontend-backend/component-library.md`
- `aiDoc/examples/README.md`
- `aiDoc/examples/backend/`
- `aiDoc/examples/frontend/`
- `aiDoc/examples/plugin/`
- `aiDoc/memory/README.md`
- `aiDoc/memory/project-memory.md`
- `aiDoc/memory/long-term/`
- `aiDoc/memory/business/`
