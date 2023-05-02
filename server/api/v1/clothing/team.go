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

type TeamApi struct {
}

var teamService = service.ServiceGroupApp.ClothingServiceGroup.TeamService


// CreateTeam 创建Team
// @Tags Team
// @Summary 创建Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Team true "创建Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/createTeam [post]
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

// DeleteTeam 删除Team
// @Tags Team
// @Summary 删除Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Team true "删除Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /team/deleteTeam [delete]
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

// DeleteTeamByIds 批量删除Team
// @Tags Team
// @Summary 批量删除Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /team/deleteTeamByIds [delete]
func (teamApi *TeamApi) DeleteTeamByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := teamService.DeleteTeamByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeam 更新Team
// @Tags Team
// @Summary 更新Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Team true "更新Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /team/updateTeam [put]
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

// FindTeam 用id查询Team
// @Tags Team
// @Summary 用id查询Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Team true "用id查询Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /team/findTeam [get]
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
