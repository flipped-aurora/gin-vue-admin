import service from '@/utils/request'

// @Tags NestExecRecord
// @Summary 创建NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestExecRecord true "创建NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtERecord/createNestExecRecord [post]
export const createNestExecRecord = (data) => {
  return service({
    url: '/NtERecord/createNestExecRecord',
    method: 'post',
    data
  })
}

// @Tags NestExecRecord
// @Summary 删除NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestExecRecord true "删除NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NtERecord/deleteNestExecRecord [delete]
export const deleteNestExecRecord = (data) => {
  return service({
    url: '/NtERecord/deleteNestExecRecord',
    method: 'delete',
    data
  })
}

// @Tags NestExecRecord
// @Summary 删除NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NtERecord/deleteNestExecRecord [delete]
export const deleteNestExecRecordByIds = (data) => {
  return service({
    url: '/NtERecord/deleteNestExecRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags NestExecRecord
// @Summary 更新NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestExecRecord true "更新NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NtERecord/updateNestExecRecord [put]
export const updateNestExecRecord = (data) => {
  return service({
    url: '/NtERecord/updateNestExecRecord',
    method: 'put',
    data
  })
}

// @Tags NestExecRecord
// @Summary 用id查询NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NestExecRecord true "用id查询NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NtERecord/findNestExecRecord [get]
export const findNestExecRecord = (params) => {
  return service({
    url: '/NtERecord/findNestExecRecord',
    method: 'get',
    params
  })
}

// @Tags NestExecRecord
// @Summary 分页获取NestExecRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取NestExecRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtERecord/getNestExecRecordList [get]
export const getNestExecRecordList = (params) => {
  return service({
    url: '/NtERecord/getNestExecRecordList',
    method: 'get',
    params
  })
}
