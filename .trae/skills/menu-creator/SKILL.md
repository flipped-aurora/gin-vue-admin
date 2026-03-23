---
name: menu-creator
description: 菜单创建工具说明
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# 菜单创建器 (create_menu)

## 功能概述

在数据库中创建前端菜单记录。

## 重要限制

**当使用 `gva_execute` 且 `needCreatedModules=true` 时，菜单会自动生成，不需要调用此工具！**

仅在以下场景使用：

1. 单独创建菜单（不涉及模块创建）
2. AI 编辑器自动添加前端页面

## 参数说明

| 参数 | 必填 | 说明 |
|------|------|------|
| `path` | 是 | 路由 path，如 `userList` |
| `name` | 是 | 路由 name，用于 Vue Router |
| `component` | 是 | Vue 组件路径，如 `view/user/list.vue` |
| `title` | 是 | 菜单显示标题 |
| `parentId` | 否 | 父菜单 ID，0 表示根菜单 |
| `icon` | 否 | 菜单图标，默认 `menu` |
| `sort` | 否 | 排序号，数字越小越靠前 |
| `hidden` | 否 | 是否隐藏 |
| `keepAlive` | 否 | 是否缓存页面 |

## 调用示例

```json
{
  "path": "productList",
  "name": "productList",
  "component": "view/product/list.vue",
  "title": "商品列表",
  "parentId": 0,
  "icon": "goods",
  "sort": 1
}
```

## 菜单按钮

```json
{
  "menuBtn": "[{\"name\":\"add\",\"desc\":\"新增\"},{\"name\":\"delete\",\"desc\":\"删除\"}]"
}
```

## 路由参数

```json
{
  "parameters": "[{\"type\":\"params\",\"key\":\"id\",\"value\":\"1\"}]"
}
```

