package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CinemaStatisticsApi struct {
}

var cinemaStatisticsService = service.ServiceGroupApp.CinemaServiceGroup.CinemaStatisticsService

// GetCinemaStatisticsList 分页获取cinemaStatistics表列表
// @Tags CinemaStatistics
// @Summary 分页获取cinemaStatistics表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cinemaReq.CinemaStatisticsSearch true "分页获取cinemaStatistics表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaStatistics/getCinemaStatisticsList [get]
func (cinemaStatisticsApi *CinemaStatisticsApi) GetCinemaStatisticsList(c *gin.Context) {
	var pageInfo cinemaReq.CinemaStatisticsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cinemaStatisticsService.GetCinemaStatisticsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
