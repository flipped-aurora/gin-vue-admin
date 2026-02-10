---
name: api-creator
description: API权限创建工具说明
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# API 权限创建器 (create_api)

## 功能概述

在数据库中创建 API 权限记录，用于 RBAC 权限控制。

## 重要限制

**当使用 `gva_execute` 且 `needCreatedModules=true` 时，API 会自动生成，不需要调用此工具！**

仅在以下场景使用：

1. 单独创建 API（不涉及模块创建）
2. AI 编辑器自动添加 API 接口
3. router 下的文件产生路径变化

## 参数说明

| 参数 | 必填 | 说明 |
|------|------|------|
| `path` | 是 | API 路径，如 `/user/create` |
| `description` | 是 | API 中文描述，如 `创建用户` |
| `apiGroup` | 是 | API 分组，如 `用户管理` |
| `method` | 否 | HTTP 方法，默认 `POST` |
| `apis` | 否 | 批量创建 JSON，如 `[{"path":"/user/list","description":"用户列表","apiGroup":"用户管理","method":"GET"}]` |

## 调用示例

```json
{
  "path": "/product/create",
  "description": "创建商品",
  "apiGroup": "商品管理",
  "method": "POST"
}
```

## 批量创建

```json
{
  "apis": "[{\"path\":\"/product/create\",\"description\":\"创建商品\",\"apiGroup\":\"商品管理\",\"method\":\"POST\"},{\"path\":\"/product/list\",\"description\":\"商品列表\",\"apiGroup\":\"商品管理\",\"method\":\"GET\"}]"
}
```

