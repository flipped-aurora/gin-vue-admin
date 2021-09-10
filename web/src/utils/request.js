import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { store } from '@/store'
import { emitter } from '@/utils/bus.js'

const service = axios.create({
    baseURL: import.meta.env.VITE_BASE_API,
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
            emitter.emit('showLoading')
        }
    }, 400)
}

const closeLoading = () => {
        acitveAxios--
        if (acitveAxios <= 0) {
            clearTimeout(timer)
            emitter.emit('closeLoading')
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
    response => {
        closeLoading()
        if (response.headers['new-token']) {
            store.commit('user/setToken', response.headers['new-token'])
        }
        if (response.data.code === 0 || response.headers.success === 'true') {
            if (response.headers.msg) {
                response.data.msg = decodeURI(response.headers.msg)
            }
            return response.data
        } else {
            ElMessage({
                showClose: true,
                message: response.data.msg,
                type: 'error'
            })
            if (response.data.data && response.data.data.reload) {
                store.commit('user/LoginOut')
            }
            return response.data.msg ? response.data : response
        }
    },
    error => {
        closeLoading()
        switch (error.response.status) {
            case 500:
                ElMessageBox.confirm(`
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `, '接口报错', {
                        dangerouslyUseHTMLString: true,
                        distinguishCancelAndClose: true,
                        confirmButtonText: '清理缓存',
                        cancelButtonText: '取消'
                    })
                    .then(() => {
                        store.commit('user/LoginOut')
                    })
                break
            case 404:
                ElMessageBox.confirm(`
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `, '接口报错', {
                    dangerouslyUseHTMLString: true,
                    distinguishCancelAndClose: true,
                    confirmButtonText: '我知道了',
                    cancelButtonText: '取消'
                })
                break
        }

        return error
    }
)

export default service