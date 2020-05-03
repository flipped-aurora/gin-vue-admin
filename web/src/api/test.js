import service from '@/utils/request'

// @Tags TestTest
// @Summary 创建TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "创建TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/createTestTest [post]
export const createTestTest = (data) => {
     return service({
         url: "/t/createTestTest",
         method: 'post',
         data
     })
 }


// @Tags TestTest
// @Summary 删除TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "删除TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /t/deleteTestTest [delete]
 export const deleteTestTest = (data) => {
     return service({
         url: "/t/deleteTestTest",
         method: 'delete',
         data
     })
 }

// @Tags TestTest
// @Summary 更新TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "更新TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /t/updateTestTest [put]
 export const updateTestTest = (data) => {
     return service({
         url: "/t/updateTestTest",
         method: 'put',
         data
     })
 }


// @Tags TestTest
// @Summary 用id查询TestTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTest true "用id查询TestTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /t/findTestTest [get]
 export const findTestTest = (params) => {
     return service({
         url: "/t/findTestTest",
         method: 'get',
         params
     })
 }


// @Tags TestTest
// @Summary 分页获取TestTest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TestTest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/getTestTestList [get]
 export const getTestTestList = (params) => {
     return service({
         url: "/t/getTestTestList",
         method: 'get',
         params
     })
 }