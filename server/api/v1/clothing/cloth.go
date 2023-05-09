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

type ClothApi struct {
}

var clothService = service.ServiceGroupApp.ClothingServiceGroup.ClothService

// CreateCloth 创建Cloth
// @Tags Cloth
// @Summary 创建Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Cloth true "创建Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloth/createCloth [post]
func (clothApi *ClothApi) CreateCloth(c *gin.Context) {
	var cloth clothing.Cloth
	err := c.ShouldBindJSON(&cloth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cloth.CreatedBy = utils.GetUserID(c)
	if !userRoleService.CheckManager(cloth.CreatedBy, cloth.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := clothService.CreateCloth(&cloth); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCloth 删除Cloth
// @Tags Cloth
// @Summary 删除Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Cloth true "删除Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloth/deleteCloth [delete]
func (clothApi *ClothApi) DeleteCloth(c *gin.Context) {
	var cloth clothing.Cloth
	err := c.ShouldBindJSON(&cloth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cloth.DeletedBy = utils.GetUserID(c)
	if !userRoleService.CheckManager(cloth.CreatedBy, cloth.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := clothService.DeleteCloth(cloth); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (clothApi *ClothApi) DeleteClothByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := clothService.DeleteClothByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCloth 更新Cloth
// @Tags Cloth
// @Summary 更新Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Cloth true "更新Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cloth/updateCloth [put]
func (clothApi *ClothApi) UpdateCloth(c *gin.Context) {
	var cloth clothing.Cloth
	err := c.ShouldBindJSON(&cloth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cloth.UpdatedBy = utils.GetUserID(c)
	if !userRoleService.CheckManager(cloth.CreatedBy, cloth.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := clothService.UpdateCloth(cloth); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCloth 用id查询Cloth
// @Tags Cloth
// @Summary 用id查询Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Cloth true "用id查询Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cloth/findCloth [get]
func (clothApi *ClothApi) FindCloth(c *gin.Context) {
	var cloth clothing.Cloth
	err := c.ShouldBindQuery(&cloth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recloth, err := clothService.GetCloth(cloth.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recloth": recloth}, c)
	}
}

// GetClothList 分页获取Cloth列表
// @Tags Cloth
// @Summary 分页获取Cloth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.ClothSearch true "分页获取Cloth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloth/getClothList [get]
func (clothApi *ClothApi) GetClothList(c *gin.Context) {
	var pageInfo clothingReq.ClothSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := clothService.GetClothInfoList(pageInfo); err != nil {
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
