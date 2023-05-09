package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TeamApplyApi struct {
}

var teamApplyService = service.ServiceGroupApp.ClothingServiceGroup.TeamApplyService

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

func (teamApplyApi *TeamApplyApi) DeleteTeamApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := teamApplyService.DeleteTeamApplyByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

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

// OptApply 审核申请
// @Tags TeamApply
// @Summary 审核申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.OptCompanyApply true "审核申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamApply/optApply [put]
func (teamApplyApi *TeamApplyApi) OptApply(c *gin.Context) {
	var opt clothingReq.OptTeamApply
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	apply, err := teamApplyService.GetTeamApply(uint(opt.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	team, err := teamService.GetTeam(apply.TeamID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if team.UserID != userID {
		global.GVA_LOG.Error("权限不足!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := teamApplyService.OptApply(apply, opt.Status, team); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
