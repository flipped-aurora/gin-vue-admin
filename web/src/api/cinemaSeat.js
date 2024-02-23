import service from '@/utils/request'

// @Tags CinemaSeat
// @Summary 创建cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaSeat true "创建cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaSeat/createCinemaSeat [post]
export const createCinemaSeat = (data) => {
  return service({
    url: '/cinemaSeat/createCinemaSeat',
    method: 'post',
    data
  })
}

// @Tags CinemaSeat
// @Summary 删除cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaSeat true "删除cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaSeat/deleteCinemaSeat [delete]
export const deleteCinemaSeat = (params) => {
  return service({
    url: '/cinemaSeat/deleteCinemaSeat',
    method: 'delete',
    params
  })
}

// @Tags CinemaSeat
// @Summary 批量删除cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaSeat/deleteCinemaSeat [delete]
export const deleteCinemaSeatByIds = (params) => {
  return service({
    url: '/cinemaSeat/deleteCinemaSeatByIds',
    method: 'delete',
    params
  })
}

// @Tags CinemaSeat
// @Summary 更新cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaSeat true "更新cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cinemaSeat/updateCinemaSeat [put]
export const updateCinemaSeat = (data) => {
  return service({
    url: '/cinemaSeat/updateCinemaSeat',
    method: 'put',
    data
  })
}

// @Tags CinemaSeat
// @Summary 用id查询cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CinemaSeat true "用id查询cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cinemaSeat/findCinemaSeat [get]
export const findCinemaSeat = (params) => {
  return service({
    url: '/cinemaSeat/findCinemaSeat',
    method: 'get',
    params
  })
}

// @Tags CinemaSeat
// @Summary 分页获取cinemaSeat表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取cinemaSeat表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaSeat/getCinemaSeatList [get]
export const getCinemaSeatList = (params) => {
  return service({
    url: '/cinemaSeat/getCinemaSeatList',
    method: 'get',
    params
  })
}
