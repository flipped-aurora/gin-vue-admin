# 子插件安装前后端契约

## 上传请求

`POST /autoCode/installPlugin` 沿用已有的 `plug` ZIP 文件字段，并支持可选的 multipart 表单字段 `parentPlugin`。

- `parentPlugin` 为空时，保持原有的顶级插件安装行为。
- 安装页面顶部继续用于安装顶级插件；每个已安装的全栈插件行提供“安装子插件”操作。打开该行的上传弹窗时，将该插件名作为 `parentPlugin` 提交。
- 后端会校验压缩包中每个前端或后端组件对应的父插件目录是否存在。

## 安装目录

传入 `parentPlugin` 时，名为 `<child>` 的子插件安装到以下两个位置：

```text
server/plugin/<parent>/subPlugin/<child>/
web/src/plugin/<parent>/subPlugin/<child>/
```

后端会在复制前端或后端任一侧文件前校验父插件目录。

## 启动与前端页面解析

安装过程中，子插件后端包会成为父插件包的一部分：

- 内部 Go import 会改为嵌套后的包路径。
- 移除子插件自身的顶级 v2 注册调用。
- 父插件会 import 子插件，并在自身的 `Register` 方法中调用 `child.Plugin.Register(group)`。
- 不修改框架的顶级插件注册方法，只调整已安装子插件的源码和父插件源码。

处理 `initialize/menu.go` 时，会规范化菜单组件路径：

- 旧路径 `plugin/<child>/...` 会变为 `plugin/<parent>/subPlugin/<child>/...`。
- 已经带有 `plugin/<parent>/subPlugin/<child>/...` 前缀的路径保持不变。

该路径与前端异步路由的插件 glob 一致，路由会相对于 `web/src/plugin` 解析组件。

## 前端源码迁移

前端文件复制到父插件目录前，安装器会扫描 Vue、JavaScript 和 TypeScript 源文件，将子插件内部别名：

```text
@/plugin/<child>/...
```

改为：

```text
@/plugin/<parent>/subPlugin/<child>/...
```

同时支持以 `<child>` 结尾的根别名。替换按路径段判断，不会修改名称相近的其他插件；已经使用父插件嵌套前缀的别名也会保持不变。
