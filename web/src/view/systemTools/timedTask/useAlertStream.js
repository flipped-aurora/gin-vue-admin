import { onMounted, onBeforeUnmount } from 'vue'
import { ElNotification } from 'element-plus'

// 订阅定时任务失败告警(SSE)。鉴权走同源 cookie x-token
// (服务端 utils/claims.go GetToken 支持 cookie 回退, EventSource 无需自定义 header)。
export const useTimedTaskAlert = () => {
  let es = null
  let failures = 0

  const connect = () => {
    const base = import.meta.env.VITE_BASE_API
    es = new EventSource(`${base}/timedTask/alertStream`)
    es.addEventListener('timedTask:alert', (e) => {
      try {
        const data = JSON.parse(e.data)
        ElNotification({
          title: `定时任务失败：${data.name}`,
          message: data.error,
          type: 'error',
          duration: 10000
        })
      } catch {
        /* 忽略非法负载 */
      }
    })
    es.onopen = () => {
      failures = 0
    }
    es.onerror = () => {
      failures += 1
      // 无权限(403)或持续失败: 关闭并放弃, 避免无限重连刷请求
      if (failures >= 3 && es) {
        es.close()
        es = null
      }
    }
  }

  onMounted(connect)
  onBeforeUnmount(() => {
    if (es) es.close()
  })
}
