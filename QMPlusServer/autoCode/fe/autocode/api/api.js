import service from '@/utils/request'

// @Tags Test
// @Summary 创建Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "创建Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/createTest [post]
export const createTest = (data) => {
     return service({
         url: "/t/createTest",
         method: 'post',
         data
     })
 }


// @Tags Test
// @Summary 删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /t/deleteTest [post]
 export const deleteTest = (data) => {
     return service({
         url: "/t/deleteTest",
         method: 'post',
         data
     })
 }

// @Tags Test
// @Summary 更新Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "更新Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /t/updateTest [post]
 export const updateTest = (data) => {
     return service({
         url: "/t/updateTest",
         method: 'post',
         data
     })
 }


// @Tags Test
// @Summary 用id查询Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "用id查询Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /t/findTest [post]
 export const findTest = (data) => {
     return service({
         url: "/t/findTest",
         method: 'post',
         data
     })
 }


// @Tags Test
// @Summary 分页获取Test列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Test列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/getTestList [post]
 export const getTestList = (data) => {
     return service({
         url: "/t/getTestList",
         method: 'post',
         data
     })
 }