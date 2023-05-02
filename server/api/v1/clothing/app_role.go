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

type AppRoleApi struct {
}

var appRoleService = service.ServiceGroupApp.ClothingServiceGroup.AppRoleService


// CreateAppRole 创建AppRole
// @Tags AppRole
// @Summary 创建AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.AppRole true "创建AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRole/createAppRole [post]
func (appRoleApi *AppRoleApi) CreateAppRole(c *gin.Context) {
	var appRole clothing.AppRole
	err := c.ShouldBindJSON(&appRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    appRole.CreatedBy = utils.GetUserID(c)
	if err := appRoleService.CreateAppRole(&appRole); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppRole 删除AppRole
// @Tags AppRole
// @Summary 删除AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.AppRole true "删除AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRole/deleteAppRole [delete]
func (appRoleApi *AppRoleApi) DeleteAppRole(c *gin.Context) {
	var appRole clothing.AppRole
	err := c.ShouldBindJSON(&appRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    appRole.DeletedBy = utils.GetUserID(c)
	if err := appRoleService.DeleteAppRole(appRole); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppRoleByIds 批量删除AppRole
// @Tags AppRole
// @Summary 批量删除AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appRole/deleteAppRoleByIds [delete]
func (appRoleApi *AppRoleApi) DeleteAppRoleByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := appRoleService.DeleteAppRoleByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppRole 更新AppRole
// @Tags AppRole
// @Summary 更新AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.AppRole true "更新AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appRole/updateAppRole [put]
func (appRoleApi *AppRoleApi) UpdateAppRole(c *gin.Context) {
	var appRole clothing.AppRole
	err := c.ShouldBindJSON(&appRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    appRole.UpdatedBy = utils.GetUserID(c)
	if err := appRoleService.UpdateAppRole(appRole); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppRole 用id查询AppRole
// @Tags AppRole
// @Summary 用id查询AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.AppRole true "用id查询AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appRole/findAppRole [get]
func (appRoleApi *AppRoleApi) FindAppRole(c *gin.Context) {
	var appRole clothing.AppRole
	err := c.ShouldBindQuery(&appRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappRole, err := appRoleService.GetAppRole(appRole.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappRole": reappRole}, c)
	}
}

// GetAppRoleList 分页获取AppRole列表
// @Tags AppRole
// @Summary 分页获取AppRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.AppRoleSearch true "分页获取AppRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRole/getAppRoleList [get]
func (appRoleApi *AppRoleApi) GetAppRoleList(c *gin.Context) {
	var pageInfo clothingReq.AppRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appRoleService.GetAppRoleInfoList(pageInfo); err != nil {
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
