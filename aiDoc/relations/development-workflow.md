# 开发工作流

## GVA Helper / MCP 约束

如果当前环境可用 GVA Helper 或其他项目专用 MCP 工具，开发前应优先使用它获取项目级建议、约束和示例，再落地具体实现。

## 推荐开发顺序

1. 先分析需求与接口
2. 先设计后端模型和请求结构
3. 再实现 Service 层业务逻辑
4. 再实现 API 层与 Router 层
5. 最后补齐 `initialize/`、插件入口或前端接入
6. 完成后进行联调与验证

## 前后端协作顺序

- 后端优先给出稳定接口
- 前端可基于 Mock 或 Swagger 并行开发
- 联调时以后端真实接口契约为准

## 分支策略

- `main`: 生产分支
- `develop`: 开发分支
- `feature/*`: 功能分支
- `hotfix/*`: 热修复分支

## 提交规范

建议使用语义化提交：

- `feat`
- `fix`
- `docs`
- `style`
- `refactor`
- `test`
- `chore`

推荐格式：`type(scope): description`
