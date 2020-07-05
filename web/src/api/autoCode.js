import service from '@/utils/request'
// @Tags api
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const createTemp = (data) => {
    return service({
        url: "/autoCode/createTemp",
        method: 'post',
        data,
        responseType: 'blob'
    })
}


// @Tags SysApi
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getDatabase [get]
export const getDB = () => {
    return service({
        url: "/autoCode/getDB",
        method: 'get',
    })
}



// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getTables [get]
export const getTable = (params) => {
    return service({
        url: "/autoCode/getTables",
        method: 'get',
        params,
    })
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getColume [get]
export const getColume = (params) => {
    return service({
        url: "/autoCode/getColume",
        method: 'get',
        params,
    })
}