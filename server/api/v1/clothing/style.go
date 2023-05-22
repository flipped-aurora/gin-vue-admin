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

type StyleApi struct {
}

var styleService = service.ServiceGroupApp.ClothingServiceGroup.StyleService

// CreateStyle 创建Style
// @Tags Style
// @Summary 创建Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Style true "创建Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /style/createStyle [post]
func (styleApi *StyleApi) CreateStyle(c *gin.Context) {
	var style clothing.Style
	err := c.ShouldBindJSON(&style)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	style.CreatedBy = utils.GetUserID(c)
	if !userRoleService.CheckManager(utils.GetUserID(c), style.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := styleService.CreateStyle(&style); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStyle 删除Style
// @Tags Style
// @Summary 删除Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Style true "删除Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /style/deleteStyle [delete]
func (styleApi *StyleApi) DeleteStyle(c *gin.Context) {
	var style clothing.Style
	err := c.ShouldBindJSON(&style)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	style.CreatedBy = utils.GetUserID(c)
	if !userRoleService.CheckManager(utils.GetUserID(c), style.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := styleService.DeleteStyle(style); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (styleApi *StyleApi) DeleteStyleByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := styleService.DeleteStyleByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStyle 更新Style
// @Tags Style
// @Summary 更新Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Style true "更新Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /style/updateStyle [put]
func (styleApi *StyleApi) UpdateStyle(c *gin.Context) {
	var style clothing.Style
	err := c.ShouldBindJSON(&style)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	style.UpdatedBy = utils.GetUserID(c)
	if !userRoleService.CheckManager(utils.GetUserID(c), style.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := styleService.UpdateStyle(style); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStyle 用id查询Style
// @Tags Style
// @Summary 用id查询Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Style true "用id查询Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /style/findStyle [get]
func (styleApi *StyleApi) FindStyle(c *gin.Context) {
	var style clothing.Style
	err := c.ShouldBindQuery(&style)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if restyle, err := styleService.GetStyle(style.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restyle": restyle}, c)
	}
}

// GetStyleList 分页获取Style列表
// @Tags Style
// @Summary 分页获取Style列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.StyleSearch true "分页获取Style列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /style/getStyleList [get]
func (styleApi *StyleApi) GetStyleList(c *gin.Context) {
	var pageInfo clothingReq.StyleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := styleService.GetStyleInfoList(pageInfo); err != nil {
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
