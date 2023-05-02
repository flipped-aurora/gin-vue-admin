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

type UserRoleApi struct {
}

var userRoleService = service.ServiceGroupApp.ClothingServiceGroup.UserRoleService


// CreateUserRole 创建UserRole
// @Tags UserRole
// @Summary 创建UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.UserRole true "创建UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userRole/createUserRole [post]
func (userRoleApi *UserRoleApi) CreateUserRole(c *gin.Context) {
	var userRole clothing.UserRole
	err := c.ShouldBindJSON(&userRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userRole.CreatedBy = utils.GetUserID(c)
	if err := userRoleService.CreateUserRole(&userRole); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserRole 删除UserRole
// @Tags UserRole
// @Summary 删除UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.UserRole true "删除UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userRole/deleteUserRole [delete]
func (userRoleApi *UserRoleApi) DeleteUserRole(c *gin.Context) {
	var userRole clothing.UserRole
	err := c.ShouldBindJSON(&userRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userRole.DeletedBy = utils.GetUserID(c)
	if err := userRoleService.DeleteUserRole(userRole); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserRoleByIds 批量删除UserRole
// @Tags UserRole
// @Summary 批量删除UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userRole/deleteUserRoleByIds [delete]
func (userRoleApi *UserRoleApi) DeleteUserRoleByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := userRoleService.DeleteUserRoleByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserRole 更新UserRole
// @Tags UserRole
// @Summary 更新UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.UserRole true "更新UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userRole/updateUserRole [put]
func (userRoleApi *UserRoleApi) UpdateUserRole(c *gin.Context) {
	var userRole clothing.UserRole
	err := c.ShouldBindJSON(&userRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userRole.UpdatedBy = utils.GetUserID(c)
	if err := userRoleService.UpdateUserRole(userRole); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserRole 用id查询UserRole
// @Tags UserRole
// @Summary 用id查询UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.UserRole true "用id查询UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userRole/findUserRole [get]
func (userRoleApi *UserRoleApi) FindUserRole(c *gin.Context) {
	var userRole clothing.UserRole
	err := c.ShouldBindQuery(&userRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserRole, err := userRoleService.GetUserRole(userRole.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserRole": reuserRole}, c)
	}
}

// GetUserRoleList 分页获取UserRole列表
// @Tags UserRole
// @Summary 分页获取UserRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.UserRoleSearch true "分页获取UserRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userRole/getUserRoleList [get]
func (userRoleApi *UserRoleApi) GetUserRoleList(c *gin.Context) {
	var pageInfo clothingReq.UserRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userRoleService.GetUserRoleInfoList(pageInfo); err != nil {
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
