import service from '@/utils/request'

// @Tags EqtInfo
// @Summary 创建EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EqtInfo true "创建EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eqtInfo/createEqtInfo [post]
export const createEqtInfo = (data) => {
  return service({
    url: '/eqtInfo/createEqtInfo',
    method: 'post',
    data
  })
}

// @Tags EqtInfo
// @Summary 删除EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EqtInfo true "删除EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eqtInfo/deleteEqtInfo [delete]
export const deleteEqtInfo = (data) => {
  return service({
    url: '/eqtInfo/deleteEqtInfo',
    method: 'delete',
    data
  })
}

// @Tags EqtInfo
// @Summary 删除EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eqtInfo/deleteEqtInfo [delete]
export const deleteEqtInfoByIds = (data) => {
  return service({
    url: '/eqtInfo/deleteEqtInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags EqtInfo
// @Summary 更新EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EqtInfo true "更新EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eqtInfo/updateEqtInfo [put]
export const updateEqtInfo = (data) => {
  return service({
    url: '/eqtInfo/updateEqtInfo',
    method: 'put',
    data
  })
}

// @Tags EqtInfo
// @Summary 用id查询EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EqtInfo true "用id查询EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eqtInfo/findEqtInfo [get]
export const findEqtInfo = (params) => {
  return service({
    url: '/eqtInfo/findEqtInfo',
    method: 'get',
    params
  })
}

// @Tags EqtInfo
// @Summary 分页获取EqtInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EqtInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eqtInfo/getEqtInfoList [get]
export const getEqtInfoList = (params) => {
  return service({
    url: '/eqtInfo/getEqtInfoList',
    method: 'get',
    params
  })
}
