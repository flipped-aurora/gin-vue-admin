package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CinemaOrderApi struct {
}

var cinemaOrderService = service.ServiceGroupApp.CinemaServiceGroup.CinemaOrderService

// FindCinemaOrder 用id查询cinemaOrder表
// @Tags CinemaOrder
// @Summary 用id查询cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cinema.CinemaOrder true "用id查询cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cinemaOrder/findCinemaOrder [get]
func (cinemaOrderApi *CinemaOrderApi) FindCinemaOrder(c *gin.Context) {
	ID := c.Query("ID")
	if recinemaOrder, err := cinemaOrderService.GetCinemaOrder(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recinemaOrder": recinemaOrder}, c)
	}
}

// GetCinemaOrderList 分页获取cinemaOrder表列表
// @Tags CinemaOrder
// @Summary 分页获取cinemaOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cinemaReq.CinemaOrderSearch true "分页获取cinemaOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaOrder/getCinemaOrderList [get]
func (cinemaOrderApi *CinemaOrderApi) GetCinemaOrderList(c *gin.Context) {
	var pageInfo cinemaReq.CinemaOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cinemaOrderService.GetCinemaOrderInfoList(pageInfo); err != nil {
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
