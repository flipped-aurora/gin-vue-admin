# 将 6 个 AI 能力迁移到独立 ai 插件

## 基本信息

- 提出日期：2026-06-24
- 当前状态：`done`
- 需求类型：插件能力迁移 / 重构（不改业务逻辑）
- 优先级：高
- 需求文件：`aiDoc/memory/business/done/ai-plugin-migration.md`

## 用户原始意图摘要

把现有分散在 `auto` 插件与 core 里的若干 AI 能力，整体「搬」到新的独立 `server/plugin/ai` 插件下，**保持逻辑不变**。

## 影响范围

- 后端：`server/plugin/ai/`（新）、`server/plugin/auto/`（剥离）、`server/source/system/menu.go`（菜单种子）
- 前端：`web/src/plugin/ai/`（新）、`web/src/plugin/auto/`（剥离）、`web/src/pathInfo.json`
- 文档：本记忆 + `demand-index.md`
- 插件 / 模块：从 `auto` 拆分出 `ai`

## 涉及对象

- 迁移的 6 个菜单能力（新建一级菜单「AI」，Name `ai`，6 个能力挂其下）：
  - `mcpTool`（Mcp Tools模板）、`mcpTest`（Mcp Tools管理）= MCP tools 生成与管理
  - `skills`（Skills管理）
  - `cli`（AI CLI管理）
  - `picture`（AI页面绘制）= 「AI 绘图」
  - `aiWorkflow`（AI 工作流）

## 已确认约束

- 采用「重新分组」而非「完整抽离」：
  - **CLI** 自包含，Go 全栈（api/service/router/model + 测试）物理迁移到 `server/plugin/ai/`
  - **skills / mcp / aiWorkflow / picture** 的 handler/service 仍留在 core（`server/api/v1/system`、`server/service/system`），`ai` 插件只搬路由 wiring、API 注册、菜单、前端视图
  - URL 路径、handler 逻辑保持不变（mcp/aiWorkflow 路由仍挂 `/autoCode/*`；picture 仍用 `/autoCode/llmAuto`）
  - API 注册 `ApiGroup` 字段保持原值，casbin 基于路径不变

## 当前进展

- ✅ 后端：CLI 全栈 + aiWorkflow 的 AutoMigrate model 已迁到 `server/plugin/ai/`，导入路径改写
- ✅ 后端：`ai` 插件 glue 完整（plugin.go / enter.go×3 / initialize 的 api·menu·router·gorm / skills·mcp·aiWorkflow 路由）
- ✅ 后端：从 `auto` 剥离已迁移项，删除 `auto/router/sys_skills.go`
- ✅ 前端：6 个视图 + cli/skills 前端 api 迁到 `web/src/plugin/ai/`，cli 视图 import 改写，pathInfo.json 更新
- ✅ 菜单种子 `source/system/menu.go`：新增「AI」父级、6 个子菜单改 Component 路径并改挂父级
- ✅ 校验：`go build`（改动包 exit 0）、`go test ./plugin/ai/...`（service ok）、前端 `vite build`（exit 0）

## 后续待办 / 注意

- ⚠️ 既有数据库：插件 `RegisterMenus` / 种子 `FirstOrCreate` 均按 Name 去重，对**已存在**的菜单行不会改写 ParentId/Component。已运行过的库需要重新初始化或在菜单管理里手动调整，新装/重置库则直接生效（这是项目既有的种子机制，非本次引入）
- `server/plugin/announcement/gen/gen.go:5` 的 “misplaced compiler directive” 是历史问题，与本次迁移无关
