# 前端 API 示例

## 这个文件负责什么

前端 API 文件负责把后端接口封装成可复用函数，统一走 `@/utils/request`，不在组件里直接拼 axios 请求。

## 什么时候应该这样写

- 新增模块接口
- 给页面提供列表、详情、创建、更新、删除方法
- 为插件页面补接口封装

## 推荐写法示例

```js
import service from '@/utils/request'

// @Summary 分页获取订单列表
// @Router /order/getOrderList [post]
export const getOrderList = (data) => {
  return service({
    url: '/order/getOrderList',
    method: 'post',
    data
  })
}

// @Summary 创建订单
// @Router /order/createOrder [post]
export const createOrder = (data) => {
  return service({
    url: '/order/createOrder',
    method: 'post',
    data
  })
}
```

## 为什么这样写

- 所有请求统一走 `service`，自动复用 token、loading、错误处理
- 页面层只调用函数，不关心底层请求细节
- JSDoc 风格的接口说明能帮助 AI 和协作者快速理解用途

## 常见错误

- 在页面组件里直接写 axios
- 把页面状态逻辑混进 API 文件
- URL、method、参数位置写错，导致接口契约漂移

## 真实参考文件

- `web/src/api/user.js`
- `web/src/plugin/announcement/api/info.js`
