---
name: gva-review
description: 代码审查工具说明
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# 代码审查器 (gva_review)

## 功能概述

在 `gva_execute` 执行后审查生成的代码，验证是否满足原始需求。

## 审查内容

1. **功能完整性**：是否覆盖所有需求功能
2. **关联正确性**：模块间关联关系是否正确实现
3. **交互功能**：用户交互和业务流程是否完整
4. **缺失识别**：发现遗漏的功能点

## 输入参数

| 参数 | 说明 |
|------|------|
| `userRequirement` | 原始需求描述 |
| `generatedFiles` | gva_execute 生成的文件列表(JSON) |

## 输出内容

- 审查结果（通过/需调整）
- 缺失功能分析
- 调整建议
- 新的开发 prompt

## 调用时机

```
gva_execute → gva_review → (手动调整)
```

## 审查重点

- 模块间关联是否正确
- CRUD 操作是否完整
- 搜索/筛选功能是否满足需求
- 字典选项是否合理

