import service from '@/utils/request'

// GetCinemaStatisticsToday 获取今日数据
// @Tags GetCinemaStatisticsToday
// @Summary 获取今日数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaStatistics/getCinemaStatisticsToday [get]
export const getCinemaStatisticsToday = () => {
  return service({
    url: '/cinemaStatistics/getCinemaStatisticsToday',
    method: 'get',
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
