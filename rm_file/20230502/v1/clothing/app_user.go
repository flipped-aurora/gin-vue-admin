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

type AppUserApi struct {
}

var appUserService = service.ServiceGroupApp.ClothingServiceGroup.AppUserService


// CreateAppUser 创建AppUser
// @Tags AppUser
// @Summary 创建AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.AppUser true "创建AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUser/createAppUser [post]
func (appUserApi *AppUserApi) CreateAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    appUser.CreatedBy = utils.GetUserID(c)
	if err := appUserService.CreateAppUser(&appUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppUser 删除AppUser
// @Tags AppUser
// @Summary 删除AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.AppUser true "删除AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUser/deleteAppUser [delete]
func (appUserApi *AppUserApi) DeleteAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    appUser.DeletedBy = utils.GetUserID(c)
	if err := appUserService.DeleteAppUser(appUser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppUserByIds 批量删除AppUser
// @Tags AppUser
// @Summary 批量删除AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appUser/deleteAppUserByIds [delete]
func (appUserApi *AppUserApi) DeleteAppUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := appUserService.DeleteAppUserByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppUser 更新AppUser
// @Tags AppUser
// @Summary 更新AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.AppUser true "更新AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUser/updateAppUser [put]
func (appUserApi *AppUserApi) UpdateAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    appUser.UpdatedBy = utils.GetUserID(c)
	if err := appUserService.UpdateAppUser(appUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppUser 用id查询AppUser
// @Tags AppUser
// @Summary 用id查询AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.AppUser true "用id查询AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUser/findAppUser [get]
func (appUserApi *AppUserApi) FindAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindQuery(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappUser, err := appUserService.GetAppUser(appUser.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappUser": reappUser}, c)
	}
}

// GetAppUserList 分页获取AppUser列表
// @Tags AppUser
// @Summary 分页获取AppUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.AppUserSearch true "分页获取AppUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUser/getAppUserList [get]
func (appUserApi *AppUserApi) GetAppUserList(c *gin.Context) {
	var pageInfo clothingReq.AppUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUserService.GetAppUserInfoList(pageInfo); err != nil {
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
