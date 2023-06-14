package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TeamApi struct {
}

var teamService = service.ServiceGroupApp.ClothingServiceGroup.TeamService

func (teamApi *TeamApi) CreateTeam(c *gin.Context) {
	var team clothing.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	team.CreatedBy = utils.GetUserID(c)
	if err := teamService.CreateTeam(&team); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (teamApi *TeamApi) DeleteTeam(c *gin.Context) {
	var team clothing.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	team.DeletedBy = utils.GetUserID(c)
	if err := teamService.DeleteTeam(team); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (teamApi *TeamApi) DeleteTeamByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := teamService.DeleteTeamByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (teamApi *TeamApi) UpdateTeam(c *gin.Context) {
	var team clothing.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	team.UpdatedBy = utils.GetUserID(c)
	if err := teamService.UpdateTeam(team); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (teamApi *TeamApi) FindTeam(c *gin.Context) {
	var team clothing.Team
	err := c.ShouldBindQuery(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteam, err := teamService.GetTeam(team.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteam": reteam}, c)
	}
}

// GetTeamList 分页获取Team列表
// @Tags Team
// @Summary 分页获取Team列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.TeamSearch true "分页获取Team列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/getTeamList [get]
func (teamApi *TeamApi) GetTeamList(c *gin.Context) {
	var pageInfo clothingReq.TeamSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teamService.GetTeamInfoList(pageInfo); err != nil {
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

// JoinTeam 组长拉用户进组
// @Tags Team
// @Summary 组长拉用户进组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.JoinTeam true "组长拉用户进组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/joinTeam [post]
func (teamApi *TeamApi) JoinTeam(c *gin.Context) {
	var req clothingReq.JoinTeam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	team, err := teamService.GetTeam(req.TeamID)
	if err != nil {
		response.FailWithMessage("组不存在", c)
		return
	}
	if team.UserID != userID {
		response.FailWithMessage("权限不足", c)
		return
	}
	company, err := companyService.GetCompany(team.CompanyID)
	if err != nil {
		response.FailWithMessage("公司不存在", c)
		return
	}
	if _, err := appUserService.GetAppUser(req.UserID); err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	if err := companyApplyService.JoinCompany(enum.Worker, req.UserID, 0, company); err != nil {
		global.GVA_LOG.Sugar().Error(err)
		response.FailWithMessage("加入公司失败", c)
		return
	}
	if err := teamService.JoinTeam(req.UserID, team); err != nil {
		global.GVA_LOG.Sugar().Error(err)
		response.FailWithMessage("加入组失败", c)
		return
	}
	response.Ok(c)
}

func (teamApi *TeamApi) DeleteMember(c *gin.Context) {
	var req clothingReq.JoinTeam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	team, err := teamService.GetTeam(req.TeamID)
	if err != nil {
		response.FailWithMessage("组不存在", c)
		return
	}
	if team.UserID != userID {
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := teamService.DeleteTeamMember(team, req.UserID); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}
