import service from '@/utils/request'

// @Tags CinemaStatistics
// @Summary 创建cinemaStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaStatistics true "创建cinemaStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaStatistics/createCinemaStatistics [post]
export const createCinemaStatistics = (data) => {
  return service({
    url: '/cinemaStatistics/createCinemaStatistics',
    method: 'post',
    data
  })
}

// @Tags CinemaStatistics
// @Summary 删除cinemaStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaStatistics true "删除cinemaStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaStatistics/deleteCinemaStatistics [delete]
export const deleteCinemaStatistics = (params) => {
  return service({
    url: '/cinemaStatistics/deleteCinemaStatistics',
    method: 'delete',
    params
  })
}

// @Tags CinemaStatistics
// @Summary 批量删除cinemaStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除cinemaStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaStatistics/deleteCinemaStatistics [delete]
export const deleteCinemaStatisticsByIds = (params) => {
  return service({
    url: '/cinemaStatistics/deleteCinemaStatisticsByIds',
    method: 'delete',
    params
  })
}

// @Tags CinemaStatistics
// @Summary 更新cinemaStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaStatistics true "更新cinemaStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cinemaStatistics/updateCinemaStatistics [put]
export const updateCinemaStatistics = (data) => {
  return service({
    url: '/cinemaStatistics/updateCinemaStatistics',
    method: 'put',
    data
  })
}

// @Tags CinemaStatistics
// @Summary 用id查询cinemaStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CinemaStatistics true "用id查询cinemaStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cinemaStatistics/findCinemaStatistics [get]
export const findCinemaStatistics = (params) => {
  return service({
    url: '/cinemaStatistics/findCinemaStatistics',
    method: 'get',
    params
  })
}

// @Tags CinemaStatistics
// @Summary 分页获取cinemaStatistics表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取cinemaStatistics表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaStatistics/getCinemaStatisticsList [get]
export const getCinemaStatisticsList = (params) => {
  return service({
    url: '/cinemaStatistics/getCinemaStatisticsList',
    method: 'get',
    params
  })
}
