import service from '@/utils/request'

// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/createSysOperationRecord [post]
export const createSysOperationRecord = (data) => {
     return service({
         url: "/sysOperationRecord/createSysOperationRecord",
         method: 'post',
         data
     })
 }


// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
 export const deleteSysOperationRecord = (data) => {
     return service({
         url: "/sysOperationRecord/deleteSysOperationRecord",
         method: 'delete',
         data
     })
 }

// @Tags SysOperationRecord
// @Summary 更新SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "更新SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysOperationRecord/updateSysOperationRecord [put]
 export const updateSysOperationRecord = (data) => {
     return service({
         url: "/sysOperationRecord/updateSysOperationRecord",
         method: 'put',
         data
     })
 }


// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "用id查询SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecord/findSysOperationRecord [get]
 export const findSysOperationRecord = (params) => {
     return service({
         url: "/sysOperationRecord/findSysOperationRecord",
         method: 'get',
         params
     })
 }


// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysOperationRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
 export const getSysOperationRecordList = (params) => {
     return service({
         url: "/sysOperationRecord/getSysOperationRecordList",
         method: 'get',
         params
     })
 }