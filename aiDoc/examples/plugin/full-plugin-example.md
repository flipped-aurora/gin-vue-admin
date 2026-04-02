# 完整插件示例

## 这个文件负责什么

这个示例用于说明：当一个能力需要以插件方式存在时，前后端目录、初始化入口和接口组织应该如何落位。

## 推荐目录结构

```text
server/plugin/order/
├── api/
│   ├── enter.go
│   └── order.go
├── config/
│   └── config.go
├── initialize/
│   ├── api.go
│   ├── gorm.go
│   ├── menu.go
│   ├── router.go
│   └── viper.go
├── model/
│   ├── order.go
│   └── request/order.go
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
├── components/
├── form/
├── view/
│   └── index.vue
└── config.js
```

## 推荐实现顺序

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

## 真实参考文件

- `server/plugin/announcement/plugin.go`
- `server/plugin/announcement/initialize/router.go`
- `server/plugin/announcement/api/enter.go`
- `server/plugin/announcement/service/enter.go`
- `web/src/plugin/announcement/api/info.js`
