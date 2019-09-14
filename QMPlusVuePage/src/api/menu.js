import service from '@/utils/request'

// @Summary 用户登录
// @Produce  application/json
// @Param 可以什么都不填 调一下即可
// @Router /menu/getMenu [post]
export const asyncMenu = () => {
    return service({
        url: "/menu/getMenu",
        method: 'post',
    })
}