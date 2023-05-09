package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CompanyApplyApi struct {
}

var companyApplyService = service.ServiceGroupApp.ClothingServiceGroup.CompanyApplyService

func (companyApplyApi *CompanyApplyApi) CreateCompanyApply(c *gin.Context) {
	// var companyApply clothing.CompanyApply
	// err := c.ShouldBindJSON(&companyApply)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	// companyApply.CreatedBy = utils.GetUserID(c)
	// if err := companyApplyService.CreateCompanyApply(&companyApply); err != nil {
	//     global.GVA_LOG.Error("创建失败!", zap.Error(err))
	// 	response.FailWithMessage("创建失败", c)
	// } else {
	// 	response.OkWithMessage("创建成功", c)
	// }
}

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

func (companyApplyApi *CompanyApplyApi) DeleteCompanyApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := companyApplyService.DeleteCompanyApplyByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

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
	userID := utils.GetUserID(c)
	if recompanyApply, err := companyApplyService.GetCompanyApply(companyApply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		if userID != recompanyApply.Company.UserID {
			response.FailWithMessage("权限不足", c)
			return
		}
		response.OkWithData(gin.H{"recompanyApply": recompanyApply}, c)
	}
}

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

// OptApply 审核申请
// @Tags CompanyApply
// @Summary 审核申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.OptCompanyApply true "审核申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyApply/optApply [put]
func (companyApplyApi *CompanyApplyApi) OptApply(c *gin.Context) {
	var opt clothingReq.OptCompanyApply
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	apply, err := companyApplyService.GetCompanyApply(uint(opt.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	company, err := companyService.GetCompany(apply.CompanyID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if company.UserID != userID {
		global.GVA_LOG.Error("权限不足!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if apply.RoleID == enum.GroupLeader {
		if _, err := teamService.GetTeamByName(company, apply.Remark); err == nil {
			opt.Status = enum.ApplyReject
		}
	}
	if err := companyApplyService.OptApply(apply, opt.Status, company); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
