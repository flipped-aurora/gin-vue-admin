import service from '@/utils/request'

// @Tags CroppingRecord
// @Summary 创建CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CroppingRecord true "创建CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /croppingRecord/createCroppingRecord [post]
export const createCroppingRecord = (data) => {
  return service({
    url: '/croppingRecord/createCroppingRecord',
    method: 'post',
    data
  })
}

// @Tags CroppingRecord
// @Summary 删除CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CroppingRecord true "删除CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /croppingRecord/deleteCroppingRecord [delete]
export const deleteCroppingRecord = (data) => {
  return service({
    url: '/croppingRecord/deleteCroppingRecord',
    method: 'delete',
    data
  })
}

// @Tags CroppingRecord
// @Summary 删除CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /croppingRecord/deleteCroppingRecord [delete]
export const deleteCroppingRecordByIds = (data) => {
  return service({
    url: '/croppingRecord/deleteCroppingRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags CroppingRecord
// @Summary 更新CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CroppingRecord true "更新CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /croppingRecord/updateCroppingRecord [put]
export const updateCroppingRecord = (data) => {
  return service({
    url: '/croppingRecord/updateCroppingRecord',
    method: 'put',
    data
  })
}

// @Tags CroppingRecord
// @Summary 用id查询CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CroppingRecord true "用id查询CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /croppingRecord/findCroppingRecord [get]
export const findCroppingRecord = (params) => {
  return service({
    url: '/croppingRecord/findCroppingRecord',
    method: 'get',
    params
  })
}

// @Tags CroppingRecord
// @Summary 分页获取CroppingRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CroppingRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /croppingRecord/getCroppingRecordList [get]
export const getCroppingRecordList = (params) => {
  return service({
    url: '/croppingRecord/getCroppingRecordList',
    method: 'get',
    params
  })
}
