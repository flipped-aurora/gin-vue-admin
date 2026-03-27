import axios from 'axios'
import { useUserStore } from '@/pinia/modules/user'
import { ElLoading, ElMessage } from 'element-plus'
import { emitter } from '@/utils/bus'
import router from '@/router/index'

const DEFAULT_REQUEST_TIMEOUT = 1000 * 60 * 10
const DEFAULT_LOADING_FORCE_CLOSE_DELAY = 30000

const service = axios.create()

let activeAxios = 0
let persistentLoadingCount = 0
let timer = null
let forceCloseTimer = null
let loadingInstance = null
let isLoadingVisible = false

const clearLoadingTimers = () => {
  if (timer) {
    clearTimeout(timer)
    timer = null
  }

  if (forceCloseTimer) {
    clearTimeout(forceCloseTimer)
    forceCloseTimer = null
  }
}

const closeLoadingInstance = () => {
  if (isLoadingVisible && loadingInstance) {
    loadingInstance.close()
  }
  loadingInstance = null
  isLoadingVisible = false
}

const scheduleForceClose = () => {
  if (!isLoadingVisible || activeAxios <= 0 || persistentLoadingCount > 0) {
    return
  }

  forceCloseTimer = setTimeout(() => {
    if (isLoadingVisible && loadingInstance) {
      console.warn(
        `Loading force closed after ${DEFAULT_LOADING_FORCE_CLOSE_DELAY}ms`
      )
      closeLoadingInstance()
      activeAxios = 0
      persistentLoadingCount = 0
    }
  }, DEFAULT_LOADING_FORCE_CLOSE_DELAY)
}

const showLoading = (
  option = {
    target: null
  }
) => {
  const loadDom = document.getElementById('gva-base-load-dom')
  const loadingOption = {
    target: null,
    ...option
  }
  const persistLoading = Boolean(loadingOption.persistLoading)

  delete loadingOption.persistLoading

  activeAxios++
  if (persistLoading) {
    persistentLoadingCount++
  }

  clearLoadingTimers()

  timer = setTimeout(() => {
    if (activeAxios > 0 && !isLoadingVisible) {
      if (!loadingOption.target) {
        loadingOption.target = loadDom
      }
      loadingInstance = ElLoading.service(loadingOption)
      isLoadingVisible = true
    }

    scheduleForceClose()
  }, 400)
}

const closeLoading = (option = {}) => {
  activeAxios--
  if (option?.persistLoading && persistentLoadingCount > 0) {
    persistentLoadingCount--
  }

  if (activeAxios <= 0) {
    activeAxios = 0
    persistentLoadingCount = 0
    clearLoadingTimers()
    closeLoadingInstance()
    return
  }

  if (forceCloseTimer) {
    clearTimeout(forceCloseTimer)
    forceCloseTimer = null
  }

  scheduleForceClose()
}

const resetLoading = () => {
  activeAxios = 0
  persistentLoadingCount = 0
  clearLoadingTimers()
  closeLoadingInstance()
}

service.interceptors.request.use(
  (config) => {
    if (typeof config.timeout === 'undefined') {
      config.timeout = DEFAULT_REQUEST_TIMEOUT
    }

    if (!config.donNotShowLoading) {
      showLoading(config.loadingOption)
    }

    config.baseURL = config.baseURL || import.meta.env.VITE_BASE_API

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
    if (!error.config?.donNotShowLoading) {
      closeLoading(error.config?.loadingOption)
    }

    emitter.emit('show-error', {
      code: 'request',
      message: error.message || '请求发送失败'
    })

    return error
  }
)

function getErrorMessage(error) {
  return error.response?.data?.msg || error.response?.statusText || '请求失败'
}

service.interceptors.response.use(
  (response) => {
    const userStore = useUserStore()

    if (!response.config.donNotShowLoading) {
      closeLoading(response.config.loadingOption)
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
    }

    ElMessage({
      showClose: true,
      message: response.data.msg || decodeURI(response.headers.msg),
      type: 'error'
    })

    return response.data.msg ? response.data : response
  },
  (error) => {
    if (!error.config?.donNotShowLoading) {
      closeLoading(error.config?.loadingOption)
    }

    if (!error.response) {
      resetLoading()
      emitter.emit('show-error', {
        code: 'network',
        message: getErrorMessage(error)
      })
      return Promise.reject(error)
    }

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

if (typeof window !== 'undefined') {
  window.addEventListener('beforeunload', resetLoading)
  window.addEventListener('unload', resetLoading)
}

export { resetLoading }
export default service
