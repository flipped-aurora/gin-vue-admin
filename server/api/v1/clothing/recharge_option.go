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

type RechargeOptionApi struct {
}

var rechargeOptionService = service.ServiceGroupApp.ClothingServiceGroup.RechargeOptionService


// CreateRechargeOption 创建RechargeOption
// @Tags RechargeOption
// @Summary 创建RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.RechargeOption true "创建RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rechargeOption/createRechargeOption [post]
func (rechargeOptionApi *RechargeOptionApi) CreateRechargeOption(c *gin.Context) {
	var rechargeOption clothing.RechargeOption
	err := c.ShouldBindJSON(&rechargeOption)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    rechargeOption.CreatedBy = utils.GetUserID(c)
	if err := rechargeOptionService.CreateRechargeOption(&rechargeOption); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRechargeOption 删除RechargeOption
// @Tags RechargeOption
// @Summary 删除RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.RechargeOption true "删除RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rechargeOption/deleteRechargeOption [delete]
func (rechargeOptionApi *RechargeOptionApi) DeleteRechargeOption(c *gin.Context) {
	var rechargeOption clothing.RechargeOption
	err := c.ShouldBindJSON(&rechargeOption)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    rechargeOption.DeletedBy = utils.GetUserID(c)
	if err := rechargeOptionService.DeleteRechargeOption(rechargeOption); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRechargeOptionByIds 批量删除RechargeOption
// @Tags RechargeOption
// @Summary 批量删除RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rechargeOption/deleteRechargeOptionByIds [delete]
func (rechargeOptionApi *RechargeOptionApi) DeleteRechargeOptionByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := rechargeOptionService.DeleteRechargeOptionByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRechargeOption 更新RechargeOption
// @Tags RechargeOption
// @Summary 更新RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.RechargeOption true "更新RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rechargeOption/updateRechargeOption [put]
func (rechargeOptionApi *RechargeOptionApi) UpdateRechargeOption(c *gin.Context) {
	var rechargeOption clothing.RechargeOption
	err := c.ShouldBindJSON(&rechargeOption)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    rechargeOption.UpdatedBy = utils.GetUserID(c)
	if err := rechargeOptionService.UpdateRechargeOption(rechargeOption); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRechargeOption 用id查询RechargeOption
// @Tags RechargeOption
// @Summary 用id查询RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.RechargeOption true "用id查询RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rechargeOption/findRechargeOption [get]
func (rechargeOptionApi *RechargeOptionApi) FindRechargeOption(c *gin.Context) {
	var rechargeOption clothing.RechargeOption
	err := c.ShouldBindQuery(&rechargeOption)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerechargeOption, err := rechargeOptionService.GetRechargeOption(rechargeOption.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerechargeOption": rerechargeOption}, c)
	}
}

// GetRechargeOptionList 分页获取RechargeOption列表
// @Tags RechargeOption
// @Summary 分页获取RechargeOption列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.RechargeOptionSearch true "分页获取RechargeOption列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rechargeOption/getRechargeOptionList [get]
func (rechargeOptionApi *RechargeOptionApi) GetRechargeOptionList(c *gin.Context) {
	var pageInfo clothingReq.RechargeOptionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := rechargeOptionService.GetRechargeOptionInfoList(pageInfo); err != nil {
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
