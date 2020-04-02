import service from '@/utils/request'

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/createExaCustomer [post]
export const createExaCustomer = (data) => {
    return service({
        url: "/customer/createExaCustomer",
        method: 'post',
        data
    })
}



// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/updateExaCustomer [post]
export const updateExaCustomer = (data) => {
    return service({
        url: "/customer/updateExaCustomer",
        method: 'post',
        data
    })
}


// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/deleteExaCustomer [post]
export const deleteExaCustomer = (data) => {
    return service({
        url: "/customer/deleteExaCustomer",
        method: 'post',
        data
    })
}


// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getExaCustomer [post]
export const getExaCustomer = (data) => {
    return service({
        url: "/customer/getExaCustomer",
        method: 'post',
        data
    })
}


// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/createExaCustomer [post]
export const getExaCustomerList = (data) => {
    return service({
        url: "/customer/getExaCustomerList",
        method: 'post',
        data
    })
}