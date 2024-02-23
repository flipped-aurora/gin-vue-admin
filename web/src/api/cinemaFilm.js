import service from '@/utils/request'

// @Tags CinemaFilm
// @Summary 创建cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaFilm true "创建cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaFilm/createCinemaFilm [post]
export const createCinemaFilm = (data) => {
  return service({
    url: '/cinemaFilm/createCinemaFilm',
    method: 'post',
    data
  })
}

// @Tags CinemaFilm
// @Summary 删除cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaFilm true "删除cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaFilm/deleteCinemaFilm [delete]
export const deleteCinemaFilm = (params) => {
  return service({
    url: '/cinemaFilm/deleteCinemaFilm',
    method: 'delete',
    params
  })
}

// @Tags CinemaFilm
// @Summary 批量删除cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaFilm/deleteCinemaFilm [delete]
export const deleteCinemaFilmByIds = (params) => {
  return service({
    url: '/cinemaFilm/deleteCinemaFilmByIds',
    method: 'delete',
    params
  })
}

// @Tags CinemaFilm
// @Summary 更新cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CinemaFilm true "更新cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cinemaFilm/updateCinemaFilm [put]
export const updateCinemaFilm = (data) => {
  return service({
    url: '/cinemaFilm/updateCinemaFilm',
    method: 'put',
    data
  })
}

// @Tags CinemaFilm
// @Summary 用id查询cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CinemaFilm true "用id查询cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cinemaFilm/findCinemaFilm [get]
export const findCinemaFilm = (params) => {
  return service({
    url: '/cinemaFilm/findCinemaFilm',
    method: 'get',
    params
  })
}

// @Tags CinemaFilm
// @Summary 分页获取cinemaFilm表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取cinemaFilm表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaFilm/getCinemaFilmList [get]
export const getCinemaFilmList = (params) => {
  return service({
    url: '/cinemaFilm/getCinemaFilmList',
    method: 'get',
    params
  })
}
