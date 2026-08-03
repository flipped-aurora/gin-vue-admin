# 完整插件示例

## 这个文件负责什么

这个示例用于说明：当一个能力需要以插件方式存在时，前后端目录、初始化入口和接口组织应该如何落位。

## 标准目录结构

```text
server/plugin/order/
├── api/
│   ├── enter.go
│   └── order.go
├── config/
│   └── config.go
├── gen/
│   └── gen.go
├── initialize/
│   ├── api.go
│   ├── dictionary.go
│   ├── gorm.go
│   ├── menu.go
│   ├── router.go
│   └── viper.go
├── model/
│   ├── order.go
│   └── request/order.go
├── plugin/
│   └── plugin.go
├── router/
│   ├── enter.go
│   └── order.go
├── service/
│   ├── enter.go
│   └── order.go
└── plugin.go

web/src/plugin/order/
├── api/
│   └── order.js
├── form/
│   └── order.vue
└── view/
    └── order.vue
```

## 标准实现顺序

1. 先定义 model 和 request
2. 再写 service
3. 再写 api 和 router
4. 再补 `initialize/`
5. 再补 `plugin.go`
6. 最后接前端 `api/view/form`

## 为什么这样写

- 前后端插件结构保持对称，更容易被 AI 和开发者理解
- 初始化逻辑集中在 `initialize/`，避免入口文件过重
- `plugin.go` 只做注册与调度，不承担业务细节

## 常见错误

- 插件目录缺少 `enter.go`，导致聚合方式不统一
- 后端插件已拆层，前端插件却把所有逻辑塞进一个页面
- 插件入口和初始化逻辑耦合过深，后续难以维护

## 参考文件与适用边界

- `server/plugin/announcement/plugin.go`：仅参考 v2 自注册与初始化调度
- `server/plugin/announcement/initialize/router.go`：仅参考 public/private 分组和完整中间件链
- `server/plugin/announcement/api/enter.go`：仅参考 API 分组聚合
- `server/plugin/announcement/service/enter.go`：仅参考 Service 分组聚合
- `server/resource/plugin/server/`：后端各层具体实现以涉及文件的当前模板为准
- `server/resource/plugin/web/`：前端各层具体实现以涉及文件的当前模板为准
- `aiDoc/examples/backend/` 与 `aiDoc/examples/frontend/`：具体分层职责和写法以对应示例为准

`announcement` 是历史形成的真实插件，只能在上面明确列出的文件和职责范围内参考。不得从这些结构参考推导其整个目录都是当前规范，也不得以“现有代码能运行”为理由覆盖更高优先级的规则、模板和分层示例。v1 `email` 与内部 `plugin-tool` 不作为标准业务插件参考。
