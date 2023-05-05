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

type CompanyApplyApi struct {
}

var companyApplyService = service.ServiceGroupApp.ClothingServiceGroup.CompanyApplyService


// CreateCompanyApply 创建CompanyApply
// @Tags CompanyApply
// @Summary 创建CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.CompanyApply true "创建CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyApply/createCompanyApply [post]
func (companyApplyApi *CompanyApplyApi) CreateCompanyApply(c *gin.Context) {
	var companyApply clothing.CompanyApply
	err := c.ShouldBindJSON(&companyApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    companyApply.CreatedBy = utils.GetUserID(c)
	if err := companyApplyService.CreateCompanyApply(&companyApply); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompanyApply 删除CompanyApply
// @Tags CompanyApply
// @Summary 删除CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.CompanyApply true "删除CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyApply/deleteCompanyApply [delete]
func (companyApplyApi *CompanyApplyApi) DeleteCompanyApply(c *gin.Context) {
	var companyApply clothing.CompanyApply
	err := c.ShouldBindJSON(&companyApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    companyApply.DeletedBy = utils.GetUserID(c)
	if err := companyApplyService.DeleteCompanyApply(companyApply); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompanyApplyByIds 批量删除CompanyApply
// @Tags CompanyApply
// @Summary 批量删除CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /companyApply/deleteCompanyApplyByIds [delete]
func (companyApplyApi *CompanyApplyApi) DeleteCompanyApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := companyApplyService.DeleteCompanyApplyByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompanyApply 更新CompanyApply
// @Tags CompanyApply
// @Summary 更新CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.CompanyApply true "更新CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /companyApply/updateCompanyApply [put]
func (companyApplyApi *CompanyApplyApi) UpdateCompanyApply(c *gin.Context) {
	var companyApply clothing.CompanyApply
	err := c.ShouldBindJSON(&companyApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    companyApply.UpdatedBy = utils.GetUserID(c)
	if err := companyApplyService.UpdateCompanyApply(companyApply); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCompanyApply 用id查询CompanyApply
// @Tags CompanyApply
// @Summary 用id查询CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.CompanyApply true "用id查询CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /companyApply/findCompanyApply [get]
func (companyApplyApi *CompanyApplyApi) FindCompanyApply(c *gin.Context) {
	var companyApply clothing.CompanyApply
	err := c.ShouldBindQuery(&companyApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recompanyApply, err := companyApplyService.GetCompanyApply(companyApply.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompanyApply": recompanyApply}, c)
	}
}

// GetCompanyApplyList 分页获取CompanyApply列表
// @Tags CompanyApply
// @Summary 分页获取CompanyApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.CompanyApplySearch true "分页获取CompanyApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyApply/getCompanyApplyList [get]
func (companyApplyApi *CompanyApplyApi) GetCompanyApplyList(c *gin.Context) {
	var pageInfo clothingReq.CompanyApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := companyApplyService.GetCompanyApplyInfoList(pageInfo); err != nil {
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
