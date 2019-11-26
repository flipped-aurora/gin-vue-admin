import service from '@/utils/request'
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {authorityId uint} true "删除角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
export const createWorkFlow = (data) => {
    return service({
        url: "/workflow/createWorkFlow",
        method: 'post',
        data
    })
}