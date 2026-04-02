# Pinia 示例

## 这个文件负责什么

Pinia store 负责全局状态、异步动作和跨页面共享数据，不负责页面渲染细节。

## 什么时候应该这样写

- 用户信息、路由、字典、系统参数等共享状态
- 多页面都会用到的业务状态
- 需要统一缓存或集中副作用的场景

## 推荐写法示例

```js
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getOrderList } from '@/api/order'

export const useOrderStore = defineStore('order', () => {
  const list = ref([])
  const total = ref(0)
  const loading = ref(false)

  const hasData = computed(() => list.value.length > 0)

  const fetchList = async (params) => {
    loading.value = true
    try {
      const res = await getOrderList(params)
      if (res.code === 0) {
        list.value = res.data.list
        total.value = res.data.total
      }
      return res
    } finally {
      loading.value = false
    }
  }

  const reset = () => {
    list.value = []
    total.value = 0
  }

  return {
    list,
    total,
    loading,
    hasData,
    fetchList,
    reset
  }
})
```

## 为什么这样写

- `ref + computed + async action` 是当前仓库里很自然的组织方式
- 页面层拿 store 结果即可，不必重复写状态管理逻辑
- 把 loading 和 reset 一并收进 store，调用侧更干净

## 常见错误

- 把所有局部页面状态都塞进全局 store
- 在 store 里写大量 DOM 操作
- 不做 loading / reset 管理，导致页面状态混乱

## 真实参考文件

- `web/src/pinia/modules/user.js`
- `web/src/pinia/modules/router.js`
- `web/src/pinia/modules/dictionary.js`
