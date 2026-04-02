# 系统关系图

## 根目录职责

- `server/`: 后端代码，包含路由、API、Service、Model、初始化和插件
- `web/`: 前端代码，包含页面、路由、状态、接口封装、工具函数和插件
- `deploy/`: Docker、Kubernetes 等部署相关资产
- `docs/`: 面向项目的人类文档与设计记录
- `aiDoc/`: 面向 AI 协作的结构化上下文

## 后端关系

后端保持现有分层方向：

1. `router/` 负责路由注册与中间件挂载
2. `api/` 负责参数绑定、请求校验、响应输出
3. `service/` 负责业务逻辑
4. `model/` 负责持久化模型和请求模型

`enter.go` 文件继续承担组合与暴露入口的职责。

## 前端关系

前端一般遵循以下流向：

1. `src/api/` 或 `src/plugin/<name>/api/` 负责接口调用
2. `src/pinia/` 负责共享状态
3. `src/router/` 负责路由与权限入口
4. `src/view/` 或 `src/plugin/<name>/view/` 负责页面
5. `src/utils/` 负责可复用工具函数

## 插件对称关系

如果某个能力以插件方式存在，尽量保持前后端结构对称：

- 后端：`server/plugin/<name>/`
- 前端：`web/src/plugin/<name>/`

当某个插件的职责和边界趋于稳定后，再把说明补充到 `aiDoc/modules/`。
