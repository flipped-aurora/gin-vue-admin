import service from '@/utils/request'

// @Tags ProductPlugin
// @Summary 创建ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "创建ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productPlugin/createProductPlugin [post]
export const createProductPlugin = (data) => {
     return service({
         url: "/productPlugin/createProductPlugin",
         method: 'post',
         data
     })
 }


 // @Tags ProductPlugin
// @Summary 上传File
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.File true "上传File"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productPlugin/upload [post]
export const uploadApi = (data) => {
    return service({
        url: "/productPlugin/upload",
        method: 'post',
        data
    })
}


// @Tags ProductPlugin
// @Summary 删除ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "删除ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productPlugin/deleteProductPlugin [delete]
 export const deleteProductPlugin = (data) => {
     return service({
         url: "/productPlugin/deleteProductPlugin",
         method: 'delete',
         data
     })
 }

// @Tags ProductPlugin
// @Summary 删除ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productPlugin/deleteProductPlugin [delete]
 export const deleteProductPluginByIds = (data) => {
     return service({
         url: "/productPlugin/deleteProductPluginByIds",
         method: 'delete',
         data
     })
 }

// @Tags ProductPlugin
// @Summary 更新ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "更新ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /productPlugin/updateProductPlugin [put]
 export const updateProductPlugin = (data) => {
     return service({
         url: "/productPlugin/updateProductPlugin",
         method: 'put',
         data
     })
 }


// @Tags ProductPlugin
// @Summary 用id查询ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "用id查询ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /productPlugin/findProductPlugin [get]
 export const findProductPlugin = (params) => {
     return service({
         url: "/productPlugin/findProductPlugin",
         method: 'get',
         params
     })
 }


// @Tags ProductPlugin
// @Summary 分页获取ProductPlugin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ProductPlugin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productPlugin/getProductPluginList [get]
 export const getProductPluginList = (params) => {
     return service({
         url: "/productPlugin/getProductPluginList",
         method: 'get',
         params
     })
 }