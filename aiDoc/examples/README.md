# 示例层

`aiDoc/examples/` 是讲解型示例层。

这里的文档不是要求 AI 逐字复制代码，而是告诉 AI：

- 某一层文件通常该承担什么职责
- 在这个仓库里推荐用什么组织方式
- 为什么要这样写
- 常见错误是什么
- 真正可以参考哪些现成文件

## 读取要求

### 新增后端能力时

建议按这个顺序读：

1. `backend/request-example.md`
2. `backend/model-example.md`
3. `backend/service-example.md`
4. `backend/api-example.md`
5. `backend/router-example.md`
6. `backend/enter-go-example.md`

### 新增前端能力时

建议按这个顺序读：

1. `frontend/api-example.md`
2. `frontend/pinia-example.md`
3. `frontend/view-example.md`
4. `frontend/utils-usage-example.md`

### 新增、修改或生成插件时

必须先读：

1. `plugin/full-plugin-example.md`
2. `backend/plugin-go-example.md`
3. `../modules/plugin-development.md`

还必须按涉及层读取当前 `server/resource/plugin/` 生成模板及对应后端/前端分层示例；详细门禁与检查表见 `../modules/plugin-development.md`。

## 使用原则

- 参考优先级为：`AGENTS.md` → 模块与分层规则 → 当前生成模板 → 本目录示例 → 示例明确列出的真实参考文件
- 真实代码不是天然规范；只参考示例文档明确指定的文件与职责，不随机选择现有实现整目录复制
- 示例关注结构与写法，不替代业务设计
- 如果真实代码、模板与示例不一致，按上述优先级处理，并及时更新低优先级的过时内容
