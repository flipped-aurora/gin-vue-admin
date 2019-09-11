import axios from 'axios'; // 引入axios
import { Message } from 'element-ui';
import { store } from '@/store/index'

const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 5000
})

//http request 拦截器
service.interceptors.request.use(
    config => {
        const token = store.getters['user/token']
        config.data = JSON.stringify(config.data);
        config.headers = {
            'Content-Type': 'application/json',
            'x-token': token
        }
        return config;
    },
    error => {
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
        if (response.data.success) {
            return response.data
        } else {
            Message({
                showClose: true,
                message: response.data.msg,
                type: 'error'
            })
            return Promise.reject(response.data.msg)
        }
    },
    error => {
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return Promise.reject(error)
    }
)

export default service