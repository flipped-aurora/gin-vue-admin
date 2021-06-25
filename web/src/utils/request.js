import axios from 'axios' // 引入axios
import { Message } from 'element-ui'
import { store } from '@/store'
import context from '@/main'
import { MessageBox } from 'element-ui'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 99999
})
let acitveAxios = 0
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      context.$bus.emit('showLoading')
    }
  }, 400)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    context.$bus.emit('closeLoading')
  }
}
// http request 拦截器
service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    const token = store.getters['user/token']
    const user = store.getters['user/userInfo']
    config.data = JSON.stringify(config.data)
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': token,
      'x-user-id': user.ID
    }
    return config
  },
  error => {
    closeLoading()
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  response => {
    closeLoading()

    if (response.headers['new-token']) {
      store.commit('user/setToken', response.headers['new-token'])
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      return response.data
    } else {
      Message({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: response.headers.msgtype || 'error'
      })
      if (response.data.data && response.data.data.reload) {
        store.commit('user/LoginOut')
      }
      return response.data.msg ? response.data : response
    }
  },
  error => {
    closeLoading()
    MessageBox.confirm(`
    <p>检测到接口错误${error}</p>
    <p>错误码500：此类错误内容常见于后台panic，如果影响您正常使用可强制登出清理缓存</p>
    <p>错误码404：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符</p>
    `, '接口报错', {
      dangerouslyUseHTMLString: true,
      distinguishCancelAndClose: true,
      confirmButtonText: '清理缓存',
      cancelButtonText: '取消'
    })
      .then(() => {
        store.commit('user/LoginOut')
      })
    return error
  }
)

export default service
