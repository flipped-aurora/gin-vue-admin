import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import router from '@/router/index'
import { ElLoading } from 'element-plus'

// 添加一个状态变量，用于跟踪是否已有错误弹窗显示
let errorBoxVisible = false

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
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

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

    // 如果已经有错误弹窗显示，则不再显示新的弹窗
    if (errorBoxVisible) {
      return error
    }

    if (!error.response) {
      // 网络错误时重置loading状态
      resetLoading()
      errorBoxVisible = true
      ElMessageBox.confirm(
        `
        <p>检测到请求错误</p>
        <p>${error}</p>
        `,
        '请求报错',
        {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '稍后重试',
          cancelButtonText: '取消'
        }
      ).finally(() => {
        // 弹窗关闭后重置状态
        errorBoxVisible = false
      })
      return
    }

    switch (error.response.status) {
      case 500:
        errorBoxVisible = true
        ElMessageBox.confirm(
          `
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `,
          '接口报错',
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '清理缓存',
            cancelButtonText: '取消'
          }
        ).then(() => {
          const userStore = useUserStore()
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        }).finally(() => {
          // 弹窗关闭后重置状态
          errorBoxVisible = false
        })
        break
      case 404:
        errorBoxVisible = true
        ElMessageBox.confirm(
          `
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `,
          '接口报错',
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '我知道了',
            cancelButtonText: '取消'
          }
        ).finally(() => {
          // 弹窗关闭后重置状态
          errorBoxVisible = false
        })
        break
      case 401:
        errorBoxVisible = true
        ElMessageBox.confirm(
          `
          <p>无效的令牌</p>
          <p>错误码:<span style="color:red"> 401 </span>错误信息:${error}</p>
          `,
          '身份信息',
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '重新登录',
            cancelButtonText: '取消'
          }
        ).then(() => {
          const userStore = useUserStore()
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        }).finally(() => {
          // 弹窗关闭后重置状态
          errorBoxVisible = false
        })
        break
    }

    return error
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
