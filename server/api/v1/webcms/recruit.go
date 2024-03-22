package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RecruitApi struct {
}

var recruitService = service.ServiceGroupApp.WebcmsServiceGroup.RecruitService

// CreateRecruit 创建Recruit
// @Tags Recruit
// @Summary 创建Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Recruit true "创建Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recruit/createRecruit [post]
func (recruitApi *RecruitApi) CreateRecruit(c *gin.Context) {
	var recruit webcms.Recruit
	err := c.ShouldBindJSON(&recruit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Job": {utils.NotEmpty()},
	}
	if err := utils.Verify(recruit, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMapString("siteinfo")
	recruit.Siteid = siteinfo["id"]
	if err := recruitService.CreateRecruit(&recruit); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRecruit 删除Recruit
// @Tags Recruit
// @Summary 删除Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Recruit true "删除Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recruit/deleteRecruit [delete]
func (recruitApi *RecruitApi) DeleteRecruit(c *gin.Context) {
	var recruit webcms.Recruit
	err := c.ShouldBindJSON(&recruit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := recruitService.DeleteRecruit(recruit); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRecruitByIds 批量删除Recruit
// @Tags Recruit
// @Summary 批量删除Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /recruit/deleteRecruitByIds [delete]
func (recruitApi *RecruitApi) DeleteRecruitByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := recruitService.DeleteRecruitByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRecruit 更新Recruit
// @Tags Recruit
// @Summary 更新Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Recruit true "更新Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recruit/updateRecruit [put]
func (recruitApi *RecruitApi) UpdateRecruit(c *gin.Context) {
	var recruit webcms.Recruit
	err := c.ShouldBindJSON(&recruit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Job": {utils.NotEmpty()},
	}
	if err := utils.Verify(recruit, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := recruitService.UpdateRecruit(recruit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRecruit 用id查询Recruit
// @Tags Recruit
// @Summary 用id查询Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcms.Recruit true "用id查询Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recruit/findRecruit [get]
func (recruitApi *RecruitApi) FindRecruit(c *gin.Context) {
	var recruit webcms.Recruit
	err := c.ShouldBindQuery(&recruit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerecruit, err := recruitService.GetRecruit(recruit.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerecruit": rerecruit}, c)
	}
}

// GetRecruitList 分页获取Recruit列表
// @Tags Recruit
// @Summary 分页获取Recruit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcmsReq.RecruitSearch true "分页获取Recruit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recruit/getRecruitList [get]
func (recruitApi *RecruitApi) GetRecruitList(c *gin.Context) {
	var pageInfo webcmsReq.RecruitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMapString("siteinfo")
	if list, total, err := recruitService.GetRecruitInfoList(pageInfo, siteinfo["id"]); err != nil {
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
