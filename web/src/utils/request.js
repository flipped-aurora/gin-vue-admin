import axios from 'axios' // 引入axios
import { useUserStore } from '@/pinia/modules/user'
import { ElLoading, ElMessage } from 'element-plus'
import { emitter } from '@/utils/bus'
import router from '@/router/index'

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})
let activeAxios = 0
let timer
let loadingInstance
let isLoadingVisible = false
let forceCloseTimer

const showLoading = (
  option = {
    target: null
  }
) => {
  const loadDom = document.getElementById('gva-base-load-dom')
  activeAxios++

  // 清除之前的定时器
  if (timer) {
    clearTimeout(timer)
  }

  // 清除强制关闭定时器
  if (forceCloseTimer) {
    clearTimeout(forceCloseTimer)
  }

  timer = setTimeout(() => {
    // 再次检查activeAxios状态，防止竞态条件
    if (activeAxios > 0 && !isLoadingVisible) {
      if (!option.target) option.target = loadDom
      loadingInstance = ElLoading.service(option)
      isLoadingVisible = true

      // 设置强制关闭定时器，防止loading永远不关闭（30秒超时）
      forceCloseTimer = setTimeout(() => {
        if (isLoadingVisible && loadingInstance) {
          console.warn('Loading强制关闭：超时30秒')
          loadingInstance.close()
          isLoadingVisible = false
          activeAxios = 0 // 重置计数器
        }
      }, 30000)
    }
  }, 400)
}

const closeLoading = () => {
  activeAxios--
  if (activeAxios <= 0) {
    activeAxios = 0 // 确保不会变成负数
    clearTimeout(timer)

    if (forceCloseTimer) {
      clearTimeout(forceCloseTimer)
      forceCloseTimer = null
    }

    if (isLoadingVisible && loadingInstance) {
      loadingInstance.close()
      isLoadingVisible = false
    }
    loadingInstance = null
  }
}

// 全局重置loading状态的函数，用于异常情况
const resetLoading = () => {
  activeAxios = 0
  isLoadingVisible = false

  if (timer) {
    clearTimeout(timer)
    timer = null
  }

  if (forceCloseTimer) {
    clearTimeout(forceCloseTimer)
    forceCloseTimer = null
  }

  if (loadingInstance) {
    try {
      loadingInstance.close()
    } catch (e) {
      console.warn('关闭loading时出错:', e)
    }
    loadingInstance = null
  }
}

// http request 拦截器
service.interceptors.request.use(
  (config) => {
    if (!config.donNotShowLoading) {
      showLoading(config.loadingOption)
    }
    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      ...config.headers
    }
    return config
  },
  (error) => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    emitter.emit('show-error', {
      code: 'request',
      message: error.message || '请求发送失败'
    })
    return error
  }
)

function getErrorMessage(error) {
  // 优先级： 响应体中的 msg > statusText > 默认消息
  return error.response?.data?.msg || error.response?.statusText || '请求失败'
}

// http response 拦截器
service.interceptors.response.use(
  (response) => {
    const userStore = useUserStore()
    if (!response.config.donNotShowLoading) {
      closeLoading()
    }
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (typeof response.data.code === 'undefined') {
      return response
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      return response.data.msg ? response.data : response
    }
  },
  (error) => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }

    if (!error.response) {
      // 网络错误
      resetLoading()
      emitter.emit('show-error', {
        code: 'network',
        message: getErrorMessage(error)
      })
      return Promise.reject(error)
    }

    // HTTP 状态码错误
    if (error.response.status === 401) {
      emitter.emit('show-error', {
        code: '401',
        message: getErrorMessage(error),
        fn: () => {
          const userStore = useUserStore()
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        }
      })
      return Promise.reject(error)
    }

    emitter.emit('show-error', {
      code: error.response.status,
      message: getErrorMessage(error)
    })
    return Promise.reject(error)
  }
)

// 监听页面卸载事件，确保loading被正确清理
if (typeof window !== 'undefined') {
  window.addEventListener('beforeunload', resetLoading)
  window.addEventListener('unload', resetLoading)
}

// 导出service和resetLoading函数
export { resetLoading }
export default service
