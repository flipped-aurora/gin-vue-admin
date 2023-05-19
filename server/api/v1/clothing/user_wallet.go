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

type UserWalletApi struct {
}

var userWalletService = service.ServiceGroupApp.ClothingServiceGroup.UserWalletService

// CreateUserWallet 创建UserWallet
// @Tags UserWallet
// @Summary 创建UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.UserWallet true "创建UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userWallet/createUserWallet [post]
func (userWalletApi *UserWalletApi) CreateUserWallet(c *gin.Context) {
	var userWallet clothing.UserWallet
	err := c.ShouldBindJSON(&userWallet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userWallet.CreatedBy = utils.GetUserID(c)
	if err := userWalletService.CreateUserWallet(&userWallet); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserWallet 删除UserWallet
// @Tags UserWallet
// @Summary 删除UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.UserWallet true "删除UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userWallet/deleteUserWallet [delete]
func (userWalletApi *UserWalletApi) DeleteUserWallet(c *gin.Context) {
	var userWallet clothing.UserWallet
	err := c.ShouldBindJSON(&userWallet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userWallet.DeletedBy = utils.GetUserID(c)
	if err := userWalletService.DeleteUserWallet(userWallet); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserWalletByIds 批量删除UserWallet
// @Tags UserWallet
// @Summary 批量删除UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userWallet/deleteUserWalletByIds [delete]
func (userWalletApi *UserWalletApi) DeleteUserWalletByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := userWalletService.DeleteUserWalletByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserWallet 更新UserWallet
// @Tags UserWallet
// @Summary 更新UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.UserWallet true "更新UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userWallet/updateUserWallet [put]
func (userWalletApi *UserWalletApi) UpdateUserWallet(c *gin.Context) {
	var userWallet clothing.UserWallet
	err := c.ShouldBindJSON(&userWallet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userWallet.UpdatedBy = utils.GetUserID(c)
	if err := userWalletService.UpdateUserWallet(userWallet); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserWallet 用id查询UserWallet
// @Tags UserWallet
// @Summary 用id查询UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.UserWallet true "用id查询UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userWallet/findUserWallet [get]
func (userWalletApi *UserWalletApi) FindUserWallet(c *gin.Context) {
	var userWallet clothing.UserWallet
	err := c.ShouldBindQuery(&userWallet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserWallet, err := userWalletService.GetUserWallet(userWallet.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserWallet": reuserWallet}, c)
	}
}

// GetUserWalletList 分页获取UserWallet列表
// @Tags UserWallet
// @Summary 分页获取UserWallet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.UserWalletSearch true "分页获取UserWallet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userWallet/getUserWalletList [get]
func (userWalletApi *UserWalletApi) GetUserWalletList(c *gin.Context) {
	var pageInfo clothingReq.UserWalletSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userWalletService.GetUserWalletInfoList(pageInfo); err != nil {
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

func (userWalletApi *UserWalletApi) GetMyWallet(c *gin.Context) {
	var pageInfo clothingReq.UserWalletSearch
	pageInfo.UserID = utils.GetUserID(c)
	if list, total, err := userWalletService.GetUserWalletInfoList(pageInfo); err != nil {
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
