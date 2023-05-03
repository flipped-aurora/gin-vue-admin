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

type TeamUserApi struct {
}

var teamUserService = service.ServiceGroupApp.ClothingServiceGroup.TeamUserService


// CreateTeamUser 创建TeamUser
// @Tags TeamUser
// @Summary 创建TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.TeamUser true "创建TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamUser/createTeamUser [post]
func (teamUserApi *TeamUserApi) CreateTeamUser(c *gin.Context) {
	var teamUser clothing.TeamUser
	err := c.ShouldBindJSON(&teamUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    teamUser.CreatedBy = utils.GetUserID(c)
	if err := teamUserService.CreateTeamUser(&teamUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeamUser 删除TeamUser
// @Tags TeamUser
// @Summary 删除TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.TeamUser true "删除TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamUser/deleteTeamUser [delete]
func (teamUserApi *TeamUserApi) DeleteTeamUser(c *gin.Context) {
	var teamUser clothing.TeamUser
	err := c.ShouldBindJSON(&teamUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    teamUser.DeletedBy = utils.GetUserID(c)
	if err := teamUserService.DeleteTeamUser(teamUser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamUserByIds 批量删除TeamUser
// @Tags TeamUser
// @Summary 批量删除TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teamUser/deleteTeamUserByIds [delete]
func (teamUserApi *TeamUserApi) DeleteTeamUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := teamUserService.DeleteTeamUserByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeamUser 更新TeamUser
// @Tags TeamUser
// @Summary 更新TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.TeamUser true "更新TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamUser/updateTeamUser [put]
func (teamUserApi *TeamUserApi) UpdateTeamUser(c *gin.Context) {
	var teamUser clothing.TeamUser
	err := c.ShouldBindJSON(&teamUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    teamUser.UpdatedBy = utils.GetUserID(c)
	if err := teamUserService.UpdateTeamUser(teamUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeamUser 用id查询TeamUser
// @Tags TeamUser
// @Summary 用id查询TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.TeamUser true "用id查询TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamUser/findTeamUser [get]
func (teamUserApi *TeamUserApi) FindTeamUser(c *gin.Context) {
	var teamUser clothing.TeamUser
	err := c.ShouldBindQuery(&teamUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteamUser, err := teamUserService.GetTeamUser(teamUser.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteamUser": reteamUser}, c)
	}
}

// GetTeamUserList 分页获取TeamUser列表
// @Tags TeamUser
// @Summary 分页获取TeamUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.TeamUserSearch true "分页获取TeamUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamUser/getTeamUserList [get]
func (teamUserApi *TeamUserApi) GetTeamUserList(c *gin.Context) {
	var pageInfo clothingReq.TeamUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teamUserService.GetTeamUserInfoList(pageInfo); err != nil {
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
