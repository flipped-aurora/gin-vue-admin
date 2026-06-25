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

## 类型一致性

- 同一字段在模型、请求结构、响应结构、前端使用处必须保持一致
- 状态字段、ID 字段、枚举字段、时间字段是高风险字段，必须重点检查
- 若涉及指针类型与非指针类型互转，必须在 Service 层显式处理 `nil`

## Service 层

- 只承载业务逻辑，不处理 HTTP 语义
- 不要依赖 `gin.Context`
- 函数应返回业务结果和 `error`
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
