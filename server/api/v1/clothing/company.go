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
	"strings"
)

type CompanyApi struct {
}

var companyService = service.ServiceGroupApp.ClothingServiceGroup.CompanyService

func (companyApi *CompanyApi) CreateCompany(c *gin.Context) {
	var company clothing.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	company.CreatedBy = utils.GetUserID(c)

	if _, err := appUserService.GetAppUser(company.UserID); err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	if err := companyService.CreateCompany(company.Name, company.UserID, company.AgentID); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		if strings.Contains(err.Error(), "unique") {
			response.FailWithMessage("公司名已存在", c)
			return
		}
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (companyApi *CompanyApi) DeleteCompany(c *gin.Context) {
	var company clothing.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	company.DeletedBy = utils.GetUserID(c)
	if err := companyService.DeleteCompany(company); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (companyApi *CompanyApi) DeleteCompanyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := companyService.DeleteCompanyByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (companyApi *CompanyApi) UpdateCompany(c *gin.Context) {
	var company clothing.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	company.UpdatedBy = utils.GetUserID(c)
	global.GVA_LOG.Sugar().Info(company)
	if err := companyService.UpdateCompany(company); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (companyApi *CompanyApi) FindCompany(c *gin.Context) {
	var company clothing.Company
	err := c.ShouldBindQuery(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recompany, err := companyService.GetCompany(company.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompany": recompany}, c)
	}
}

// GetCompanyList 分页获取Company列表
// @Tags Company
// @Summary 分页获取Company列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.CompanySearch true "分页获取Company列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /company/getCompanyList [get]
func (companyApi *CompanyApi) GetCompanyList(c *gin.Context) {
	var pageInfo clothingReq.CompanySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := companyService.GetCompanyInfoList(pageInfo); err != nil {
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

// JoinCompany 加入公司
// @Tags Company
// @Summary 加入公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.JoinCompany true "加入公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /company/getCompanyList [get]
func (companyApi *CompanyApi) JoinCompany(c *gin.Context) {
	var joinCompany clothingReq.JoinCompany
	err := c.ShouldBindJSON(&joinCompany)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	switch joinCompany.RoleID {
	case enum.Boss:
		if _, err := companyService.GetCompanyByName(joinCompany.Remark); err == nil {
			response.FailWithMessage("公司已存在", c)
			return
		}
		if err := companyService.CreateCompany(joinCompany.Remark, userID, 0); err != nil {
			response.FailWithMessage("服务器错误", c)
			return
		}
	case enum.Tailor:
		if company, err := companyService.GetCompany(joinCompany.CompanyID); err != nil {
			response.FailWithMessage("公司不存在", c)
			return
		} else {
			var companyApply clothing.CompanyApply
			companyApply.CompanyID = company.ID
			companyApply.UserID = userID
			companyApply.RoleID = enum.Tailor
			status := new(int)
			*status = 0
			companyApply.Status = status
			if err := companyApplyService.CreateCompanyApply(&companyApply); err != nil {
				global.GVA_LOG.Sugar().Error(err)
				response.FailWithMessage("操作失败", c)
				return
			}
			if err := msgBoxService.SendMsg(userID, company.UserID, enum.CompanyApply, companyApply.ID); err != nil {
				global.GVA_LOG.Sugar().Error(err)
				response.FailWithMessage("操作失败", c)
				return
			}
		}
	case enum.GroupLeader:
		if company, err := companyService.GetCompany(joinCompany.CompanyID); err != nil {
			response.FailWithMessage("公司不存在", c)
			return
		} else {
			if _, err := teamService.GetTeamByName(company, joinCompany.Remark); err == nil {
				response.FailWithMessage("组已存在", c)
				return
			}
			var companyApply clothing.CompanyApply
			companyApply.CompanyID = company.ID
			companyApply.UserID = userID
			companyApply.RoleID = enum.GroupLeader
			companyApply.Remark = joinCompany.Remark
			status := new(int)
			*status = 0
			companyApply.Status = status
			if err := companyApplyService.CreateCompanyApply(&companyApply); err != nil {
				global.GVA_LOG.Sugar().Error(err)
				response.FailWithMessage("操作失败", c)
				return
			}
			if err := msgBoxService.SendMsg(userID, company.UserID, enum.CompanyApply, companyApply.ID); err != nil {
				global.GVA_LOG.Sugar().Error(err)
				response.FailWithMessage("操作失败", c)
				return
			}
		}
	case enum.Worker:
		if team, err := teamService.GetTeam(joinCompany.TeamID); err != nil {
			response.FailWithMessage("组不存在", c)
			return
		} else {
			if company, err := companyService.GetCompany(team.CompanyID); err != nil {
				response.FailWithMessage("公司不存在", c)
				return
			} else {
				if err := companyApplyService.JoinCompany(enum.Worker, userID, company); err != nil {
					global.GVA_LOG.Sugar().Error(err)
					response.FailWithMessage("加入公司失败", c)
					return
				}
				var wallet clothing.UserWallet
				wallet.CompanyID = company.ID
				wallet.UserID = userID
				if err := userWalletService.CreateUserWallet(&wallet); err != nil {
					global.GVA_LOG.Sugar().Error(err)
					response.FailWithMessage("创建钱包失败", c)
					return
				}
				var teamApply clothing.TeamApply
				teamApply.TeamID = team.ID
				teamApply.UserID = userID
				status := new(int)
				*status = 0
				teamApply.Status = status
				if err := teamApplyService.CreateTeamApply(&teamApply); err != nil {
					global.GVA_LOG.Sugar().Error(err)
					response.FailWithMessage("操作失败", c)
					return
				}
				if err := msgBoxService.SendMsg(userID, team.UserID, enum.WorkerApply, teamApply.ID); err != nil {
					global.GVA_LOG.Sugar().Error(err)
					response.FailWithMessage("操作失败", c)
					return
				}
			}
		}
	default:
		response.FailWithMessage("未定义的角色类型", c)
		return
	}
	response.Ok(c)
}
