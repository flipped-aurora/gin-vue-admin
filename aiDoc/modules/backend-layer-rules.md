# 后端分层约束

## 总原则

- 严格遵守 `Router -> API -> Service -> Model` 依赖方向
- 禁止跨层直接调用
- `enter.go` 作为组装与暴露入口，避免循环引用

## Model 层

- 数据模型优先继承 `global.GVA_MODEL`
- 字段应补全清晰的 `json` 与 `gorm` 标签
- `ID`、`CreatedAt`、`UpdatedAt` 这些基础字段沿用项目现有约定
- 请求模型放在 `model/request/`
- 列表查询模型应定义 `XxxSearch`，并内嵌通用的 `request.PageInfo`
- `CreatedBy`/`UpdatedBy`/`DeletedBy`/`DeptId`（列名 `created_by`/`updated_by`/`deleted_by`/`dept_id`）这组公共操作字段**仅在业务表需要数据权限时才创建**，对应代码生成器的 AutoCreateResource 产物，不是每张表的必备字段；手写模型需要同类语义时用同名字段，不要自造 `CreatorID` 等同义字段
- 模型上的**关联对象字段**（Preload 填充的 struct / 指针，如 `User SysUser`、`Leader *SysUser`）必须加 `form:"-"`：gin 的 query/form 绑定按类型树递归且会给 nil 指针自动 new，模型间互相引用（如 `SysUser.Dept` ⇄ `SysDepartment.Leader`）一旦被 `ShouldBindQuery` 扫到会无限递归、进程直接 `stack overflow` 崩溃；关联对象本来也不可能从 query string 传入
- `DeptId`（归属部门）服务于数据权限：数据权限引擎按 `created_by`/`dept_id` 两列做行级过滤与创建时自动盖章，自造字段不会被引擎识别

## 类型一致性

- 同一字段在模型、请求结构、响应结构、前端使用处必须保持一致
- 状态字段、ID 字段、枚举字段、时间字段是高风险字段，必须重点检查
- 若涉及指针类型与非指针类型互转，必须在 Service 层显式处理 `nil`

## Service 层

- 只承载业务逻辑，不处理 HTTP 语义
- 不要依赖 `gin.Context`
- 函数应返回业务结果和 `error`
- 查询方法以 `ctx context.Context` 为首参，数据库调用用 `global.GVA_DB.WithContext(ctx)` 串联请求链路（API 层传 `c.Request.Context()`）
- 分页统一 `limit, offset := info.LimitOffset()`（`request.PageInfo` 提供，pageSize 超过 `MaxPageSize=100` 自动截断），不要手写 `PageSize*(Page-1)` 换算
- 数据权限（行级过滤）由统一引擎的 GORM 全局回调实现（`server/utils/datascope/`）：受控表（带 `dept_id`/`created_by` 列）的范围过滤与创建盖章由引擎自动完成，Service 不手写 `dept_id`/`created_by` 过滤条件、不手动赋值 `CreatedBy`/`DeptId`；更新走 `Save` 等全量写时用 `Omit("dept_id", "created_by")` 保护归属列不被表单零值覆盖
- 操作人盖章按列存在自动参与，均不手动赋值：更新盖 `updated_by`（不要放进 `Omit`）；表同时有 `deleted_by` 列与 `gorm.DeletedAt` 时，软删除的那条 UPDATE 自动并入 `deleted_by`（硬删除 / `Unscoped` 不盖）；无身份 / `WithSystem` / `UpdateColumn`(SkipHooks) 不盖
- 漏写条件的 update/delete 会被引擎挡下并报 `ErrMissingWhereClause`（不会静默作用于整个数据范围）；确需全量写用 `Session(&gorm.Session{AllowGlobalUpdate: true})` 显式声明
- 漏传 ctx 等于旁路数据权限（现阶段放行 + 告警 + 落审计表）；确需跨范围查询用 `db.Set("data_scope:skip", true)` 显式旁路，定时任务/CLI/初始化用 `datascope.WithSystem(ctx)`，不要裸用 `context.Background()`
- 每个模块在 `service/` 下建立独立文件，并在 `service/enter.go` 注册

## API 层

- 负责参数提取、参数校验、调用 Service 和统一响应
- 参数从哪里取，取决于前端怎么传、协议怎么设计、当前逻辑需要什么，以及哪个位置更合理
- 不要把绑定方式写死成某一种固定模板

### 常见参数来源

- JSON body
- Query string
- Path params
- `multipart/form-data`
- Header
- Cookie

### 常见取法

- JSON body: `ShouldBindJSON`
- Query: `ShouldBindQuery`、`c.Query(...)`、`c.DefaultQuery(...)`
- Path: `c.Param(...)`
- form-data / file upload: `c.FormFile(...)`、`c.DefaultPostForm(...)`、`c.Request.FormValue(...)`
- Header: `c.GetHeader(...)`、`c.Request.Header.Get(...)`
- Cookie: `c.Cookie(...)`

### 使用原则

- 绑定方式要与真实参数来源一致
- 不要为了套模板，把 Header / Cookie / Query / form-data 中的数据强行改成 body
- 认证、追踪、网关透传等信息，很多时候本来就应该从 Header 或 Cookie 获取
- 上传文件时，应按上传协议从 `multipart/form-data` 中取文件和附带字段

- 必须通过 `service.ServiceGroupApp` 访问服务层
- 必须使用项目统一的 `response` 包输出结果
- 每个对外 API 都必须写完整且准确的 Swagger 注释

## Router 层

- 负责路由分组、中间件挂载和处理函数绑定
- 必须通过 `api.ApiGroupApp` 引用 API 层
- 每个模块在 `router/` 下建立独立文件，并在 `router/enter.go` 注册

## Initialize 层

插件或模块若需要初始化入口，至少关注以下职责：

- `gorm.go`: 表结构迁移
- `router.go`: 路由注册
- `menu.go`: 菜单与权限初始化
- `viper.go`: 配置加载
- `api.go`: API 注册

## Swagger 约束

对外 API 的 Swagger 注释至少要准确说明：

- 功能说明
- 请求参数
- 响应结构
- 路由路径
- 鉴权要求

### 响应类型要落到具体类型

`@Success` 的 `data` 必须反映真实返回类型，让 swag 能生成有意义的返回结构，而不是空对象：

- 分页列表：`response.Response{data=response.PageResult{list=[]xxx.Model},msg=string}`
  - `response.PageResult.List` 是 `interface{}`，只写 `data=response.PageResult` 会让 swag 把 `list` 生成成空对象，必须用嵌套覆盖把元素类型补上
- 非分页列表 / 直接返回数组：`response.Response{data=[]xxx.Model,msg=string}`
- 单对象 / 详情：`response.Response{data=xxx.Model,msg=string}`
- 仅返回提示、无数据（创建 / 更新 / 删除）：`response.Response{msg=string}`
- 仅当返回的是动态结构或示例数据（数据源、临时 `gin.H` 拼装等）时，才用 `data=object` 或 `data=[]interface{}`

### 鉴权注释要与路由分组一致

- 私有分组（`PrivateGroup`，挂 `JWTAuth` + `Casbin`）的接口才写 `@Security ApiKeyAuth`
- 公开分组（`PublicGroup`）的接口不写 `@Security`，否则文档与真实鉴权不符

> 代码生成模板 `resource/package/server/api/api.go.tpl`、`resource/plugin/server/api/api.go.tpl` 已按上述规范生成列表接口返回类型，手写接口遵循同一标准。
