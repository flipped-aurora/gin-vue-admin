import axios from 'axios'; // 引入axios
import { Message, Loading } from 'element-ui';
import { store } from '@/store/index'
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 99999
})
let acitveAxios = 0
let loadingInstance
let timer
const showLoading = () => {
    acitveAxios++
    if (timer) {
        clearTimeout(timer)
    }
    timer = setTimeout(() => {
        if (acitveAxios > 0) {
            loadingInstance = Loading.service({ fullscreen: true })
        }
    }, 400);
}

const closeLoading = () => {
        acitveAxios--
        if (acitveAxios <= 0) {
            clearTimeout(timer)
            loadingInstance && loadingInstance.close()
        }
    }
    //http request 拦截器
service.interceptors.request.use(
    config => {
        showLoading()
        const token = store.getters['user/token']
        config.data = config.headers['Content-Type'] == 'application/json' ? config.data : JSON.stringify(config.data);
        config.headers = {
            'Content-Type': config.headers['Content-Type'] || 'application/json',
            'x-token': token
        }
        return config;
    },
    error => {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return Promise.reject(error);
    }
);


//http response 拦截器
service.interceptors.response.use(
    response => {
        closeLoading()
        if (response.data.success) {
            return response.data
        } else {
            Message({
                showClose: true,
                message: response.data.msg,
                type: 'error',
                onClose: () => {
                    if (response.data.data && response.data.data.reload) {
                        store.commit('user/LoginOut')
                    }
                }
            })
            return Promise.reject(response.data.msg)
        }
    },
    error => {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return Promise.reject(error)
    }
)

export default service