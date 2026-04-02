# 工具函数使用示例

## 这个文件负责什么

这个示例告诉 AI：当前端需要通用能力时，应先复用 `src/utils/`，而不是临时再造一套。

## 什么时候应该这样写

- 发送 HTTP 请求
- 格式化日期
- 获取字典数据
- 处理按钮权限
- 做命名转换
- 跨组件通信

## 推荐写法示例

```js
import service from '@/utils/request'
import { formatDate, CreateUUID } from '@/utils/format'
import { getDict } from '@/utils/dictionary'
import { useBtnAuth } from '@/utils/btnAuth'

const token = CreateUUID()
const createdAt = formatDate(new Date())

const loadStatusDict = async () => {
  return await getDict('order_status')
}

const btnAuth = useBtnAuth()

export const fetchOrderList = (data) => {
  return service({
    url: '/order/getOrderList',
    method: 'post',
    data
  })
}
```

## 为什么这样写

- 统一工具入口能减少重复实现
- `request`、`format`、`dictionary`、`btnAuth` 都已经被项目广泛使用
- 复用已有工具比新造 helper 更利于 AI 和人协作

## 常见错误

- 手写日期格式化逻辑
- 直接使用 axios 绕开 `request`
- 自己再实现一套按钮权限判断
- 明明已有字典工具，却在页面里重复请求和缓存

## 真实参考文件

- `web/src/utils/request.js`
- `web/src/utils/format.js`
- `web/src/utils/dictionary.js`
- `web/src/utils/btnAuth.js`
