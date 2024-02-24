package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CinemaSeatApi struct {
}

var cinemaSeatService = service.ServiceGroupApp.CinemaServiceGroup.CinemaSeatService

// CreateCinemaSeat 创建cinemaSeat表
// @Tags CinemaSeat
// @Summary 创建cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinemaReq.CinemaSeatCreate true "创建cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaSeat/createCinemaSeat [post]
func (cinemaSeatApi *CinemaSeatApi) CreateCinemaSeat(c *gin.Context) {
	var cinemaSeat cinemaReq.CinemaSeatCreate
	err := c.ShouldBindJSON(&cinemaSeat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cinemaSeatService.CreateCinemaSeat(&cinemaSeat); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCinemaSeat 删除cinemaSeat表
// @Tags CinemaSeat
// @Summary 删除cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaSeat true "删除cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaSeat/deleteCinemaSeat [delete]
func (cinemaSeatApi *CinemaSeatApi) DeleteCinemaSeat(c *gin.Context) {
	ID := c.Query("ID")
	if err := cinemaSeatService.DeleteCinemaSeat(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// FindCinemaSeat 用id查询cinemaSeat表
// @Tags CinemaSeat
// @Summary 用id查询cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cinema.CinemaSeat true "用id查询cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cinemaSeat/findCinemaSeat [get]
func (cinemaSeatApi *CinemaSeatApi) FindCinemaSeat(c *gin.Context) {
	ID := c.Query("ID")
	if recinemaSeat, err := cinemaSeatService.GetCinemaSeat(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recinemaSeat": recinemaSeat}, c)
	}
}

// GetCinemaSeatList 分页获取cinemaSeat表列表
// @Tags CinemaSeat
// @Summary 分页获取cinemaSeat表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cinemaReq.CinemaSeatSearch true "分页获取cinemaSeat表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaSeat/getCinemaSeatList [get]
func (cinemaSeatApi *CinemaSeatApi) GetCinemaSeatList(c *gin.Context) {
	var pageInfo cinemaReq.CinemaSeatSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cinemaSeatService.GetCinemaSeatInfoList(pageInfo); err != nil {
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
