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

type CinemaSeatApi struct {
}

var cinemaSeatService = service.ServiceGroupApp.CinemaServiceGroup.CinemaSeatService


// CreateCinemaSeat 创建cinemaSeat表
// @Tags CinemaSeat
// @Summary 创建cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaSeat true "创建cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaSeat/createCinemaSeat [post]
func (cinemaSeatApi *CinemaSeatApi) CreateCinemaSeat(c *gin.Context) {
	var cinemaSeat cinema.CinemaSeat
	err := c.ShouldBindJSON(&cinemaSeat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cinemaSeatService.CreateCinemaSeat(&cinemaSeat); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
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

// DeleteCinemaSeatByIds 批量删除cinemaSeat表
// @Tags CinemaSeat
// @Summary 批量删除cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cinemaSeat/deleteCinemaSeatByIds [delete]
func (cinemaSeatApi *CinemaSeatApi) DeleteCinemaSeatByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cinemaSeatService.DeleteCinemaSeatByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCinemaSeat 更新cinemaSeat表
// @Tags CinemaSeat
// @Summary 更新cinemaSeat表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaSeat true "更新cinemaSeat表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cinemaSeat/updateCinemaSeat [put]
func (cinemaSeatApi *CinemaSeatApi) UpdateCinemaSeat(c *gin.Context) {
	var cinemaSeat cinema.CinemaSeat
	err := c.ShouldBindJSON(&cinemaSeat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cinemaSeatService.UpdateCinemaSeat(cinemaSeat); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
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
