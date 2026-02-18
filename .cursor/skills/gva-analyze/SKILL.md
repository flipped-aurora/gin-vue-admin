---
name: gva-analyze
description: 系统分析工具说明
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# 系统分析器 (gva_analyze)

## 功能概述

分析当前系统中已有的包、模块和字典信息，判断是否需要创建新资源。

## 返回信息

### 已有包列表

```json
{
  "packages": [
    {"name": "example", "type": "package", "modules": ["User", "Product"]}
  ]
}
```

### 已有字典列表

```json
{
  "dictionaries": [
    {"type": "gender", "name": "性别", "options": ["男", "女"]}
  ]
}
```

### 分析结论

- 是否需要创建新包
- 是否需要创建新模块
- 是否需要创建新字典
- 可复用的现有资源

## 使用场景

1. 开发新功能前，了解系统现状
2. 避免重复创建已存在的资源
3. 复用现有字典和模块

## 调用时机

在 `requirement_analyzer` 之后、`gva_execute` 之前调用：

```
requirement_analyzer → gva_analyze → gva_execute
```

