package Nestrolepkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Nestrolepkg"
	nestrolepkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/Nestrolepkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NestRoleApi struct {
}

var nestroleService = service.ServiceGroupApp.NestrolepkgServiceGroup.NestRoleService

// CreateNestRole 创建NestRole
// @Tags NestRole
// @Summary 创建NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body nestrolepkg.NestRole true "创建NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestrole/createNestRole [post]
func (nestroleApi *NestRoleApi) CreateNestRole(c *gin.Context) {
	var nestrole Nestrolepkg.NestRole
	err := c.ShouldBindJSON(&nestrole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	nestrole.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Roleid": {utils.NotEmpty()},
		"Nestid": {utils.NotEmpty()},
	}
	if err := utils.Verify(nestrole, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := nestroleService.CreateNestRole(&nestrole); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNestRole 删除NestRole
// @Tags NestRole
// @Summary 删除NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Nestrolepkg.NestRole true "删除NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nestrole/deleteNestRole [delete]
func (nestroleApi *NestRoleApi) DeleteNestRole(c *gin.Context) {
	var nestrole Nestrolepkg.NestRole
	err := c.ShouldBindJSON(&nestrole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	nestrole.DeletedBy = utils.GetUserID(c)
	if err := nestroleService.DeleteNestRole(nestrole); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNestRoleByIds 批量删除NestRole
// @Tags NestRole
// @Summary 批量删除NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /nestrole/deleteNestRoleByIds [delete]
func (nestroleApi *NestRoleApi) DeleteNestRoleByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := nestroleService.DeleteNestRoleByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNestRole 更新NestRole
// @Tags NestRole
// @Summary 更新NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Nestrolepkg.NestRole true "更新NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /nestrole/updateNestRole [put]
func (nestroleApi *NestRoleApi) UpdateNestRole(c *gin.Context) {
	var nestrole Nestrolepkg.NestRole
	err := c.ShouldBindJSON(&nestrole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	nestrole.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Roleid": {utils.NotEmpty()},
		"Nestid": {utils.NotEmpty()},
	}
	if err := utils.Verify(nestrole, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := nestroleService.UpdateNestRole(nestrole); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNestRole 用id查询NestRole
// @Tags NestRole
// @Summary 用id查询NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Nestrolepkg.NestRole true "用id查询NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /nestrole/findNestRole [get]
func (nestroleApi *NestRoleApi) FindNestRole(c *gin.Context) {
	var nestrole Nestrolepkg.NestRole
	err := c.ShouldBindQuery(&nestrole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if renestrole, err := nestroleService.GetNestRole(nestrole.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renestrole": renestrole}, c)
	}
}

// GetNestRoleList 分页获取NestRole列表
// @Tags NestRole
// @Summary 分页获取NestRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query nestrolepkgReq.NestRoleSearch true "分页获取NestRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestrole/getNestRoleList [get]
func (nestroleApi *NestRoleApi) GetNestRoleList(c *gin.Context) {
	var pageInfo nestrolepkgReq.NestRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userClaims, err := utils.GetClaims(c)
	if err != nil {
		global.GVA_LOG.Error("获取userClaims失败!", zap.Error(err))
		response.FailWithMessage("获取userClaims失败", c)
		return
	}
	// user, err := service.ServiceGroupApp.SystemServiceGroup.UserService.FindUserByIdWithAuth(int(userClaims.BaseClaims.ID))
	// if err != nil {
	// 	global.GVA_LOG.Error("获取user失败!", zap.Error(err))
	// 	response.FailWithMessage("获取user失败", c)
	// 	return
	// }
	// var authIDs []uint
	// for _, item := range user.Authorities {
	// 	authIDs = append(authIDs, item.AuthorityId)
	// }
	pageInfo.AuthID, err = service.ServiceGroupApp.SystemServiceGroup.UserService.GetUserAuthorities(userClaims.BaseClaims.ID)
	if err != nil {
		global.GVA_LOG.Error("获取用户所属角色失败!", zap.Error(err))
		response.FailWithMessage("获取用户所属角色失败", c)
		return
	}

	if list, total, err := nestroleService.GetNestRoleInfoList(pageInfo); err != nil {
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
