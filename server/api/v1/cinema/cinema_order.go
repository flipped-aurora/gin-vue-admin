package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
    cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CinemaOrderApi struct {
}

var cinemaOrderService = service.ServiceGroupApp.CinemaServiceGroup.CinemaOrderService


// CreateCinemaOrder 创建cinemaOrder表
// @Tags CinemaOrder
// @Summary 创建cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaOrder true "创建cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaOrder/createCinemaOrder [post]
func (cinemaOrderApi *CinemaOrderApi) CreateCinemaOrder(c *gin.Context) {
	var cinemaOrder cinema.CinemaOrder
	err := c.ShouldBindJSON(&cinemaOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cinemaOrderService.CreateCinemaOrder(&cinemaOrder); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCinemaOrder 删除cinemaOrder表
// @Tags CinemaOrder
// @Summary 删除cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaOrder true "删除cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaOrder/deleteCinemaOrder [delete]
func (cinemaOrderApi *CinemaOrderApi) DeleteCinemaOrder(c *gin.Context) {
	ID := c.Query("ID")
	if err := cinemaOrderService.DeleteCinemaOrder(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCinemaOrderByIds 批量删除cinemaOrder表
// @Tags CinemaOrder
// @Summary 批量删除cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cinemaOrder/deleteCinemaOrderByIds [delete]
func (cinemaOrderApi *CinemaOrderApi) DeleteCinemaOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cinemaOrderService.DeleteCinemaOrderByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCinemaOrder 更新cinemaOrder表
// @Tags CinemaOrder
// @Summary 更新cinemaOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaOrder true "更新cinemaOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cinemaOrder/updateCinemaOrder [put]
func (cinemaOrderApi *CinemaOrderApi) UpdateCinemaOrder(c *gin.Context) {
	var cinemaOrder cinema.CinemaOrder
	err := c.ShouldBindJSON(&cinemaOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cinemaOrderService.UpdateCinemaOrder(cinemaOrder); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

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
