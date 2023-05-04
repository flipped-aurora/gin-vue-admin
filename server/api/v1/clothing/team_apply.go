package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type TeamApplyApi struct {
}

var teamApplyService = service.ServiceGroupApp.ClothingServiceGroup.TeamApplyService


// CreateTeamApply 创建TeamApply
// @Tags TeamApply
// @Summary 创建TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.TeamApply true "创建TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamApply/createTeamApply [post]
func (teamApplyApi *TeamApplyApi) CreateTeamApply(c *gin.Context) {
	var teamApply clothing.TeamApply
	err := c.ShouldBindJSON(&teamApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    teamApply.CreatedBy = utils.GetUserID(c)
	if err := teamApplyService.CreateTeamApply(&teamApply); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeamApply 删除TeamApply
// @Tags TeamApply
// @Summary 删除TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.TeamApply true "删除TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamApply/deleteTeamApply [delete]
func (teamApplyApi *TeamApplyApi) DeleteTeamApply(c *gin.Context) {
	var teamApply clothing.TeamApply
	err := c.ShouldBindJSON(&teamApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    teamApply.DeletedBy = utils.GetUserID(c)
	if err := teamApplyService.DeleteTeamApply(teamApply); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamApplyByIds 批量删除TeamApply
// @Tags TeamApply
// @Summary 批量删除TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teamApply/deleteTeamApplyByIds [delete]
func (teamApplyApi *TeamApplyApi) DeleteTeamApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := teamApplyService.DeleteTeamApplyByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeamApply 更新TeamApply
// @Tags TeamApply
// @Summary 更新TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.TeamApply true "更新TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamApply/updateTeamApply [put]
func (teamApplyApi *TeamApplyApi) UpdateTeamApply(c *gin.Context) {
	var teamApply clothing.TeamApply
	err := c.ShouldBindJSON(&teamApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    teamApply.UpdatedBy = utils.GetUserID(c)
	if err := teamApplyService.UpdateTeamApply(teamApply); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeamApply 用id查询TeamApply
// @Tags TeamApply
// @Summary 用id查询TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.TeamApply true "用id查询TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamApply/findTeamApply [get]
func (teamApplyApi *TeamApplyApi) FindTeamApply(c *gin.Context) {
	var teamApply clothing.TeamApply
	err := c.ShouldBindQuery(&teamApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteamApply, err := teamApplyService.GetTeamApply(teamApply.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteamApply": reteamApply}, c)
	}
}

// GetTeamApplyList 分页获取TeamApply列表
// @Tags TeamApply
// @Summary 分页获取TeamApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.TeamApplySearch true "分页获取TeamApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamApply/getTeamApplyList [get]
func (teamApplyApi *TeamApplyApi) GetTeamApplyList(c *gin.Context) {
	var pageInfo clothingReq.TeamApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teamApplyService.GetTeamApplyInfoList(pageInfo); err != nil {
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
