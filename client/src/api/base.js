import request from '@/utils/request'

export const login = (data) => {
    return request({
        url: '/login',
        method: 'post',
        data
    })
}