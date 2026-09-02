---
tool: claude
role: compatibility-adapter
canonical_source: /AGENT.MD
structured_context: /aiDoc
---

# Claude 规则适配层

本文件只用于兼容 Claude 现有的自动加载路径。

## 真实规则入口

请按下面顺序读取：

1. `/AGENT.MD`
2. `/aiDoc/README.md`
3. `/aiDoc/relations/`、`/aiDoc/modules/`、`/aiDoc/frontend-backend/`、`/aiDoc/examples/`、`/aiDoc/memory/` 中与当前任务相关的文件

## 适配层约束

- 不要在这里扩写项目级规则
- 项目级规则变更时，先更新 `/AGENT.MD` 与 `/aiDoc/`
- 工具目录只保留薄适配层职责，不再保存独立 project rule 副本
- 代码读取约束以 `/AGENT.MD` 为准，其中包括：无论什么情况，都不要直接读取 `node_modules/` 中的代码
