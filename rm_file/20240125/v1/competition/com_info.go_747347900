package competition

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/competition"
    competitionReq "github.com/flipped-aurora/gin-vue-admin/server/model/competition/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ComInfoApi struct {
}

var comDataService = service.ServiceGroupApp.CompetitionServiceGroup.ComInfoService


// CreateComInfo 创建比赛信息
// @Tags ComInfo
// @Summary 创建比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body competition.ComInfo true "创建比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /comData/createComInfo [post]
func (comDataApi *ComInfoApi) CreateComInfo(c *gin.Context) {
	var comData competition.ComInfo
	err := c.ShouldBindJSON(&comData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := comDataService.CreateComInfo(&comData); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteComInfo 删除比赛信息
// @Tags ComInfo
// @Summary 删除比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body competition.ComInfo true "删除比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comData/deleteComInfo [delete]
func (comDataApi *ComInfoApi) DeleteComInfo(c *gin.Context) {
	id := c.Query("ID")
	if err := comDataService.DeleteComInfo(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteComInfoByIds 批量删除比赛信息
// @Tags ComInfo
// @Summary 批量删除比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /comData/deleteComInfoByIds [delete]
func (comDataApi *ComInfoApi) DeleteComInfoByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := comDataService.DeleteComInfoByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateComInfo 更新比赛信息
// @Tags ComInfo
// @Summary 更新比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body competition.ComInfo true "更新比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comData/updateComInfo [put]
func (comDataApi *ComInfoApi) UpdateComInfo(c *gin.Context) {
	var comData competition.ComInfo
	err := c.ShouldBindJSON(&comData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := comDataService.UpdateComInfo(comData); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindComInfo 用id查询比赛信息
// @Tags ComInfo
// @Summary 用id查询比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query competition.ComInfo true "用id查询比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comData/findComInfo [get]
func (comDataApi *ComInfoApi) FindComInfo(c *gin.Context) {
	id := c.Query("ID")
	if recomData, err := comDataService.GetComInfo(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recomData": recomData}, c)
	}
}

// GetComInfoList 分页获取比赛信息列表
// @Tags ComInfo
// @Summary 分页获取比赛信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query competitionReq.ComInfoSearch true "分页获取比赛信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comData/getComInfoList [get]
func (comDataApi *ComInfoApi) GetComInfoList(c *gin.Context) {
	var pageInfo competitionReq.ComInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := comDataService.GetComInfoInfoList(pageInfo); err != nil {
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
