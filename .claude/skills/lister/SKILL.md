---
name: lister
description: 查询工具说明
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# 查询工具

## API 列表查询 (list_all_apis)

获取系统中所有 API 接口信息。

### 返回数据

1. **databaseApis**: 数据库中已注册的 API（包含 ID、描述、分组）
2. **ginApis**: gin 框架中实际注册的路由（仅路径和方法）

### 使用场景

- 判断是使用现有 API 还是创建新 API
- 查找可复用的 API 接口
- 前端调用时查找正确的 API 路径

### 路径示例

```shell
/api/user/:id  → 根据用户 ID 获取用户信息
/api/user/list → 获取用户列表
/api/user/create → 创建用户
```

---

## 菜单列表查询 (list_all_menus)

获取系统中所有菜单信息。

### 返回数据

- 完整的菜单树形结构
- 路由配置（path、name、component）
- 菜单元数据（title、icon、keepAlive）
- 菜单参数和按钮配置
- 父子菜单关系

### 使用场景

- 前端路由配置：配置 vue-router
- 菜单权限管理：了解可用菜单项
- 导航组件开发：构建动态导航
- 页面跳转：使用 `router.push({name: 'xxx'})`

### 菜单结构示例

```json
{
  "path": "userList",
  "name": "userList",
  "component": "view/user/list.vue",
  "meta": {
    "title": "用户列表",
    "icon": "user",
    "keepAlive": false
  },
  "children": []
}
```

---

## 前端跳转示例

根据查询到的菜单信息进行页面跳转：

```javascript
// 使用 name 跳转
router.push({ name: 'userList' })

// 带参数跳转
router.push({ name: 'userDetail', params: { id: 123 } })

// 使用 path 跳转
router.push('/user/list')
```

