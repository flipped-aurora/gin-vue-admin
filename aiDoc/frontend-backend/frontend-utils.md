# 前端工具函数约束

## 核心原则

开发任何前端功能前，必须优先检查并复用 `src/utils/` 下已有能力，严禁重复造轮子。

## 重点工具

- `request.js`: HTTP 请求统一入口
- `date.js`: 日期格式化
- `format.js`: 布尔值、字典、URL、主题色、UUID 等常用格式化能力
- `dictionary.js`: 字典数据获取
- `stringFun.js`: 命名格式转换
- `params.js`: 系统参数获取
- `bus.js`: 跨组件事件通信
- `closeThisPage.js`: 关闭当前标签页
- `downloadImg.js`: 图片下载
- `image.js`: 图片压缩
- `event.js`: DOM 事件管理
- `env.js`: 环境判断
- `doc.js`: 文档跳转
- `fmtRouterTitle.js`: 动态路由标题格式化
- `page.js`: 页面标题生成
- `asyncRouter.js`: 异步路由处理
- `btnAuth.js`: 按钮级权限控制

## 强制使用场景

- 发起 HTTP 请求时，必须使用 `@/utils/request`
- 获取字典数据时，必须优先使用 `@/utils/dictionary`
- 生成 UUID 时，必须优先使用 `CreateUUID`
- 处理按钮权限时，必须优先使用 `useBtnAuth`
- 进行命名格式转换时，必须优先使用 `@/utils/stringFun`
- 跨组件通信优先使用事件总线，避免滥用 Pinia
