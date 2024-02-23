import service from '@/utils/request'

// @Tags CinemaOrder
// @Summary 创建cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaOrder true "创建cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaOrder/createCinemaOrder [post]
export const createCinemaOrder = (data) => {
  return service({
    url: '/cinemaOrder/createCinemaOrder',
    method: 'post',
    data
  })
}

// @Tags CinemaOrder
// @Summary 删除cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaOrder true "删除cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaOrder/deleteCinemaOrder [delete]
export const deleteCinemaOrder = (params) => {
  return service({
    url: '/cinemaOrder/deleteCinemaOrder',
    method: 'delete',
    params
  })
}

// @Tags CinemaOrder
// @Summary 批量删除cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaOrder/deleteCinemaOrder [delete]
export const deleteCinemaOrderByIds = (params) => {
  return service({
    url: '/cinemaOrder/deleteCinemaOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags CinemaOrder
// @Summary 更新cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaOrder true "更新cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cinemaOrder/updateCinemaOrder [put]
export const updateCinemaOrder = (data) => {
  return service({
    url: '/cinemaOrder/updateCinemaOrder',
    method: 'put',
    data
  })
}

// @Tags CinemaOrder
// @Summary 用id查询cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CinemaOrder true "用id查询cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cinemaOrder/findCinemaOrder [get]
export const findCinemaOrder = (params) => {
  return service({
    url: '/cinemaOrder/findCinemaOrder',
    method: 'get',
    params
  })
}

// @Tags CinemaOrder
// @Summary 分页获取cinemaOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取cinemaOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaOrder/getCinemaOrderList [get]
export const getCinemaOrderList = (params) => {
  return service({
    url: '/cinemaOrder/getCinemaOrderList',
    method: 'get',
    params
  })
}
