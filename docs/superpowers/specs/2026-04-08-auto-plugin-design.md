# Auto 插件化设计

## 目标

将当前“编程辅助”整条能力线从主系统拆出，重构为可独立安装、注册、初始化、移除的 `auto` 插件。

本次设计同时满足两个约束：

- 后端接口路径保持不变，继续兼容现有 `/autoCode/*` 与 `/skills/*`
- 前端菜单位置调整为插件内统一维护，编程辅助入口不再分散挂载在主系统原位置

## 业务边界

`auto` 插件纳入以下能力：

- 自动代码 `autoCode`
- 自动化包 `autoPkg`
- 自动代码管理 `autoCodeAdmin`
- MCP 工具生成与服务管理
- AI 工作流 `AIWorkflow`
- 插件安装 `installPlugin`
- 插件打包 `pubPlug`
- Skills 管理

本次不直接迁移为插件内部实现的内容：

- 其他系统页面中对 `llmAuto` 的零散增强调用
- `server/cmd/mcp` 独立启动程序本身
- `server/mcp` 中与独立服务进程管理强绑定的底层基础设施

说明：

- 上述未迁移部分可以继续被 `auto` 插件调用，但不作为本次插件边界的一部分
- 若后续需要做到“移除 auto 插件后完全没有 MCP 相关基础设施残留”，再单独做第二轮基础设施下沉

## 设计原则

1. 插件归属单一
所有编程辅助相关业务代码、菜单、初始化逻辑、前端页面与前端 API 统一归属到 `auto` 插件目录。

2. 外部接口兼容
后端仍暴露现有接口路径，避免前端改造与客户侧联调成本扩大。

3. 拔插优先
主系统不再持有该业务的常驻注册逻辑；移除插件后，编程辅助能力整体失效，而不是残留半套功能。

4. 菜单重新归位
页面菜单按新的“编程辅助”树组织，但保留关键页面名与页面能力，降低前端功能回归风险。

## 后端结构设计

目标目录：

`server/plugin/auto/`

建议结构：

- `api/`
- `service/`
- `model/`
- `model/request/`
- `model/response/`
- `router/`
- `initialize/`
- `plugin/`
- `plugin.go`

### API 层迁移

迁入插件的 API 文件：

- `server/api/v1/system/sys_auto_code.go`
- `server/api/v1/system/sys_auto_code_sse.go`
- `server/api/v1/system/auto_code_template.go`
- `server/api/v1/system/auto_code_mcp.go`
- `server/api/v1/system/auto_code_package.go`
- `server/api/v1/system/auto_code_history.go`
- `server/api/v1/system/auto_code_plugin.go`
- `server/api/v1/system/ai_workflow_session.go`
- `server/api/v1/system/sys_skills.go`

插件内需要补齐 `api/enter.go`，形成独立的 `ApiGroup`，避免继续依赖主系统 `SystemApiGroup` 的全局聚合变量。

### Service 层迁移

迁入插件的 service 文件：

- `server/service/system/sys_auto_code_interface.go`
- `server/service/system/sys_auto_code_mysql.go`
- `server/service/system/sys_auto_code_pgsql.go`
- `server/service/system/sys_auto_code_sqlite.go`
- `server/service/system/sys_auto_code_mssql.go`
- `server/service/system/sys_auto_code_oracle.go`
- `server/service/system/auto_code_template.go`
- `server/service/system/auto_code_mcp.go`
- `server/service/system/auto_code_package.go`
- `server/service/system/auto_code_history.go`
- `server/service/system/auto_code_plugin.go`
- `server/service/system/ai_workflow_session.go`
- `server/service/system/ai_workflow_markdown.go`
- `server/service/system/sys_skills.go`

插件内需要补齐 `service/enter.go`，形成插件内部 `ServiceGroup`。

### Model 层迁移

迁入插件的 model 文件：

- `server/model/system/sys_ai_workflow_session.go`
- `server/model/system/sys_auto_code_history.go`
- `server/model/system/sys_skills.go`
- `server/model/system/request/sys_ai_workflow_session.go`
- `server/model/system/request` 中与 autoCode、skills、MCP 相关请求结构
- `server/model/system/response/sys_ai_workflow_session.go`
- `server/model/system/response` 中与 autoCode 相关响应结构

注意事项：

- 对主系统通用模型的引用可以保留，例如菜单、API、字典等系统基础模型
- 仅迁移编程辅助业务自身持有的数据模型与请求响应结构

### Router 层设计

插件内统一接管以下路由：

- `/autoCode/*`
- `/skills/*`

其中：

- `/autoCode/*` 保持现有请求方法与路径不变
- `/skills/*` 保持现有请求方法与路径不变

主系统 `server/initialize/router.go` 中将移除：

- `systemRouter.InitAutoCodeRouter(...)`
- `systemRouter.InitAutoCodeHistoryRouter(...)`
- `systemRouter.InitSkillsRouter(...)`

对应职责改由 `server/plugin/auto/initialize/router.go` 注册。

### Initialize 层设计

插件初始化至少包含：

- `initialize/router.go`
- `initialize/gorm.go`
- `initialize/menu.go`
- `initialize/api.go`
- `initialize/dictionary.go` 可选

职责划分：

- `router.go`: 注册 `/autoCode/*` 与 `/skills/*`
- `gorm.go`: 为 `SysAIWorkflowSession`、`SysAutoCodeHistory` 以及其他插件自有模型执行迁移
- `menu.go`: 注册编程辅助菜单树
- `api.go`: 注册插件菜单对应 API 权限
- `dictionary.go`: 如果编程辅助没有独立字典项，可以保留空实现

### Plugin 入口设计

插件入口实现 `Plugin` 接口，启动时执行：

1. `initialize.Api`
2. `initialize.Menu`
3. `initialize.Dictionary`
4. `initialize.Gorm`
5. `initialize.Router`

插件注册入口加入：

- `server/plugin/register.go`

主系统最终只保留插件注册，不再直接挂载 auto 业务路由。

## 前端结构设计

目标目录：

`web/src/plugin/auto/`

建议结构：

- `api/`
- `view/`
- `router/` 可选
- `menu/` 可选

### 页面迁移范围

迁入插件的页面：

- `web/src/view/systemTools/autoCode/index.vue`
- `web/src/view/systemTools/autoCode/mcp.vue`
- `web/src/view/systemTools/autoCode/mcpTest.vue`
- `web/src/view/systemTools/autoCode/picture.vue`
- `web/src/view/systemTools/autoCode/component/*`
- `web/src/view/systemTools/autoPkg/autoPkg.vue`
- `web/src/view/systemTools/autoCodeAdmin/index.vue`
- `web/src/view/systemTools/aiWrokflow/index.vue`
- `web/src/view/systemTools/installPlugin/index.vue`
- `web/src/view/systemTools/pubPlug/pubPlug.vue`
- `web/src/view/systemTools/skills/index.vue`

### 前端 API 迁移范围

迁入插件的 API 文件：

- `web/src/api/autoCode.js`
- `web/src/api/skills.js` 如果仓库中存在独立文件则一并迁移；若当前 Skills 页面实际复用其他 API 文件，则在插件中拆分归位

迁移后的要求：

- 页面内部改为引用 `@/plugin/auto/api/...`
- 对外请求路径保持不变，不修改后端接口前缀

### 前端菜单树

插件内统一维护如下菜单树：

- 编程辅助
- 自动代码
- 自动化包
- AI 工作流
- MCP 工具
- MCP 测试
- 插件安装
- 插件打包
- 技能管理

`autoCodeAdmin` 作为管理页处理：

- 方案 A：挂为“自动代码”子页面但默认隐藏
- 方案 B：不在一级菜单显示，仅通过页面跳转进入

推荐方案 B，因为其更偏编辑态页面，不适合作为常驻一级菜单。

### 路由与组件兼容

以下兼容策略保持不变：

- 页面 `name` 尽量沿用已有值，例如 `AutoCode`、`AIWorkflow`、`Skills`
- 从其他页面跳转到 `autoCode`、`autoPkg`、`AIWorkflow` 的逻辑继续可用
- 页面内部业务逻辑保持稳定，先做目录归位，再做样式与位置微调

## 主系统剥离点

需要从主系统移除的常驻业务接线：

- `server/router/system/enter.go` 中 `AutoCodeRouter`、`SkillsRouter` 的业务归属
- `server/initialize/router.go` 中对 autoCode、history、skills 的显式注册
- `server/api/v1/system/enter.go` 中对 auto 相关 service 全局变量绑定
- `server/service/system/enter.go` 中对 auto 相关 service group 聚合

需要保留的主系统能力：

- 插件注册机制
- 系统级菜单、API、字典注册基础能力
- 通用基础模型与通用中间件

## 兼容策略

### 接口兼容

保留现有路径：

- `/autoCode/getDB`
- `/autoCode/getTables`
- `/autoCode/getColumn`
- `/autoCode/preview`
- `/autoCode/createTemp`
- `/autoCode/addFunc`
- `/autoCode/mcp*`
- `/autoCode/getPackage`
- `/autoCode/createPackage`
- `/autoCode/delPackage`
- `/autoCode/getTemplates`
- `/autoCode/getMeta`
- `/autoCode/getSysHistory`
- `/autoCode/rollback`
- `/autoCode/pubPlug`
- `/autoCode/installPlugin`
- `/autoCode/removePlugin`
- `/autoCode/getPluginList`
- `/autoCode/saveAIWorkflowSession`
- `/autoCode/getAIWorkflowSessionList`
- `/autoCode/getAIWorkflowSessionDetail`
- `/autoCode/deleteAIWorkflowSession`
- `/autoCode/dumpAIWorkflowMarkdown`
- `/skills/*`

### 页面兼容

- 旧页面 `name` 尽量保持不变
- 旧跳转入口暂不大规模修改语义
- 菜单位置改变，但页面能力不变

### 基础设施兼容

- `auto` 插件可继续调用主系统现有 `server/mcp` 能力
- `cmd/mcp` 暂不迁目录，避免本轮把独立启动程序与插件化一起耦合重构

## 实施顺序

1. 建立 `server/plugin/auto` 基础目录与 `plugin.go`
2. 迁移 model/request/response
3. 迁移 service，并建立 `service/enter.go`
4. 迁移 api，并建立 `api/enter.go`
5. 迁移 router，并在插件内注册兼容路径
6. 建立 `initialize/gorm.go`
7. 建立 `initialize/menu.go`
8. 建立 `initialize/api.go`
9. 迁移前端页面到 `web/src/plugin/auto/view`
10. 迁移前端 API 到 `web/src/plugin/auto/api`
11. 调整前端菜单位置与页面引用
12. 从主系统移除 auto 业务注册点
13. 保留必要兼容引用并完成回归验证

## 测试与验证

后端验证重点：

- `/autoCode/*` 原有接口仍能访问
- `/skills/*` 原有接口仍能访问
- 插件注册后路由完整
- 插件未注册时编程辅助整体不可用
- `SysAIWorkflowSession` 与 `SysAutoCodeHistory` 自动迁移正常

前端验证重点：

- 编程辅助菜单显示位置符合新结构
- 自动代码、自动化包、AI 工作流、Skills、插件安装与打包页面可正常进入
- 原页面内部跳转仍可达
- MCP 页面状态查询与测试流程可用

## 风险与处理

风险 1：主系统聚合入口耦合过深

处理：

- 先在插件内建立完整 `enter.go` 分层，再移除主系统绑定，避免半迁移状态下引用断裂

风险 2：菜单与页面路径调整导致前端跳转失效

处理：

- 保留页面 `name`
- 先迁目录，再补菜单，再修复内部引用

风险 3：Skills 文件系统操作依赖仓库目录结构

处理：

- 保持现有工具目录解析逻辑不变
- 只迁移业务归属，不在本轮重写 Skills 文件系统根定位规则

风险 4：MCP 独立服务与插件边界不完全一致

处理：

- 本轮明确将其视为插件依赖的基础设施
- 等插件业务拆出稳定后，再判断是否继续做基础设施插件化

## 最终结果

完成后，`auto` 插件应具备以下特征：

- 编程辅助相关页面、API、服务、模型、菜单、初始化逻辑全部归属插件
- 主系统不再直接注册编程辅助业务路由
- 客户需要移除该能力时，可以以插件为边界进行拔除
- 现有接口地址保持稳定，降低已有前端和外部调用改造成本
