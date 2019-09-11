import service from '@/utils/request'

// @Summary 用户登录
// @Produce  application/json
// @Param data body {userName:"string",passWord:"string"}
// @Router /base/login [post]
export const login = (data) => {
    service({
        url: "/base/login",
        method: 'post',
        data: data
    })
}