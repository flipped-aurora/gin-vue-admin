import service from '@/utils/request'

// @Tags AerialPhotographyResult
// @Summary 创建AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AerialPhotographyResult true "创建AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ALPhotographyResult/createAerialPhotographyResult [post]
export const createAerialPhotographyResult = (data) => {
  return service({
    url: '/ALPhotographyResult/createAerialPhotographyResult',
    method: 'post',
    data
  })
}

// @Tags AerialPhotographyResult
// @Summary 删除AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AerialPhotographyResult true "删除AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ALPhotographyResult/deleteAerialPhotographyResult [delete]
export const deleteAerialPhotographyResult = (data) => {
  return service({
    url: '/ALPhotographyResult/deleteAerialPhotographyResult',
    method: 'delete',
    data
  })
}

// @Tags AerialPhotographyResult
// @Summary 删除AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ALPhotographyResult/deleteAerialPhotographyResult [delete]
export const deleteAerialPhotographyResultByIds = (data) => {
  return service({
    url: '/ALPhotographyResult/deleteAerialPhotographyResultByIds',
    method: 'delete',
    data
  })
}

// @Tags AerialPhotographyResult
// @Summary 更新AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AerialPhotographyResult true "更新AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ALPhotographyResult/updateAerialPhotographyResult [put]
export const updateAerialPhotographyResult = (data) => {
  return service({
    url: '/ALPhotographyResult/updateAerialPhotographyResult',
    method: 'put',
    data
  })
}

// @Tags AerialPhotographyResult
// @Summary 用id查询AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AerialPhotographyResult true "用id查询AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ALPhotographyResult/findAerialPhotographyResult [get]
export const findAerialPhotographyResult = (params) => {
  return service({
    url: '/ALPhotographyResult/findAerialPhotographyResult',
    method: 'get',
    params
  })
}

// @Tags AerialPhotographyResult
// @Summary 分页获取AerialPhotographyResult列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AerialPhotographyResult列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ALPhotographyResult/getAerialPhotographyResultList [get]
export const getAerialPhotographyResultList = (params) => {
  return service({
    url: '/ALPhotographyResult/getAerialPhotographyResultList',
    method: 'get',
    params
  })
}
