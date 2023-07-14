import service from '@/utils/request'

// @Tags FlyResult
// @Summary 创建FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FlyResult true "创建FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/createFlyResult [post]
export const createFlyResult = (data) => {
  return service({
    url: '/FlyRt/createFlyResult',
    method: 'post',
    data
  })
}

// @Tags FlyResult
// @Summary 删除FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FlyResult true "删除FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FlyRt/deleteFlyResult [delete]
export const deleteFlyResult = (data) => {
  return service({
    url: '/FlyRt/deleteFlyResult',
    method: 'delete',
    data
  })
}

// @Tags FlyResult
// @Summary 删除FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FlyRt/deleteFlyResult [delete]
export const deleteFlyResultByIds = (data) => {
  return service({
    url: '/FlyRt/deleteFlyResultByIds',
    method: 'delete',
    data
  })
}

// @Tags FlyResult
// @Summary 更新FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FlyResult true "更新FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FlyRt/updateFlyResult [put]
export const updateFlyResult = (data) => {
  return service({
    url: '/FlyRt/updateFlyResult',
    method: 'put',
    data
  })
}

// @Tags FlyResult
// @Summary 用id查询FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FlyResult true "用id查询FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FlyRt/findFlyResult [get]
export const findFlyResult = (params) => {
  return service({
    url: '/FlyRt/findFlyResult',
    method: 'get',
    params
  })
}

// @Tags FlyResult
// @Summary 分页获取FlyResult列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FlyResult列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/getFlyResultList [get]
export const getFlyResultList = (params) => {
  return service({
    url: '/FlyRt/getFlyResultList',
    method: 'get',
    params
  })
}
