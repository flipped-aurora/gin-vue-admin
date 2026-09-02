# 示例层

`aiDoc/examples/` 是讲解型示例层。

这里的文档不是要求 AI 逐字复制代码，而是告诉 AI：

- 某一层文件通常该承担什么职责
- 在这个仓库里推荐用什么组织方式
- 为什么要这样写
- 常见错误是什么
- 真正可以参考哪些现成文件

## 读取建议

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

### 新增插件时

建议至少先读：

1. `plugin/full-plugin-example.md`
2. `backend/plugin-go-example.md`

## 使用原则

- 优先参考仓库内真实文件，其次参考本目录的示例文档
- 示例关注结构与写法，不替代业务设计
- 如果真实代码与旧示例不一致，应以当前仓库真实结构为准，并及时更新示例
