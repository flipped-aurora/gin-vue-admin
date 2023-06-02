package clothing

import (
	"fmt"
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
	"gorm.io/gorm"
)

type JobApplyApi struct {
}

var jobApplyService = service.ServiceGroupApp.ClothingServiceGroup.JobApplyService

// CreateJobApply 创建JobApply
// @Tags JobApply
// @Summary 创建JobApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.JobApply true "创建JobApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobApply/createJobApply [post]
func (jobApplyApi *JobApplyApi) CreateJobApply(c *gin.Context) {
	var jobApply clothing.JobApply
	err := c.ShouldBindJSON(&jobApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	croppingRecord, err := croppingRecordService.GetCroppingRecord(jobApply.CroppingID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("裁剪单不存在", c)
		return
	}
	global.GVA_KEYLOCK.Lock(fmt.Sprintf("cropping-%d", croppingRecord.ID))
	defer global.GVA_KEYLOCK.Unlock(fmt.Sprintf("cropping-%d", croppingRecord.ID))
	jobApply.CreatedBy = utils.GetUserID(c)
	jobApply.UserID = utils.GetUserID(c)
	if !userRoleService.CheckStaff(jobApply.CreatedBy, croppingRecord.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	team, err := teamService.GetTeam(jobApply.TeamID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("组不存在", c)
		return
	}
	switch jobApply.JobType {
	case enum.Process:
		process, err := processService.GetProcess(jobApply.ProcessID)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("工序不存在", c)
			return
		}
		jobApply.ProcessName = process.Name
		jobApply.Price = process.Price
		var i clothing.Inventory
		global.GVA_DB.Where("cropping_record_id = ? and process_id = ? and size = ?", croppingRecord.ID, process.ID, jobApply.Size).First(&i)
		if int(i.Margin) < jobApply.Quantity {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("库存不足", c)
			return
		} else {
			global.GVA_DB.Model(&i).Update("margin", gorm.Expr("margin - ?", jobApply.Quantity))
		}
	case enum.Whole:
		style, err := styleService.GetStyle(croppingRecord.StyleID)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("款式不存在", c)
			return
		}
		jobApply.ProcessName = "成衣"
		jobApply.Price = style.Price
		var s clothing.SizeList
		global.GVA_DB.Where("cropping_record_id = ? and size = ?", croppingRecord.ID, jobApply.Size).First(&s)
		if int(s.Margin) < jobApply.Quantity {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("库存不足", c)
			return
		} else {
			global.GVA_DB.Model(&s).Update("margin", gorm.Expr("margin - ?", jobApply.Quantity))
		}
	}
	status := new(int)
	*status = 0
	jobApply.Status = status
	if err := jobApplyService.CreateJobApply(&jobApply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		err := msgBoxService.SendMsg(jobApply.CreatedBy, team.UserID, enum.JobApply, jobApply.ID)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
		response.OkWithMessage("创建成功", c)
	}
}

func (jobApplyApi *JobApplyApi) DeleteJobApply(c *gin.Context) {
	var jobApply clothing.JobApply
	err := c.ShouldBindJSON(&jobApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jobApply.DeletedBy = utils.GetUserID(c)
	if err := jobApplyService.DeleteJobApply(jobApply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (jobApplyApi *JobApplyApi) DeleteJobApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := jobApplyService.DeleteJobApplyByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (jobApplyApi *JobApplyApi) UpdateJobApply(c *gin.Context) {
	var jobApply clothing.JobApply
	err := c.ShouldBindJSON(&jobApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jobApply.UpdatedBy = utils.GetUserID(c)
	if err := jobApplyService.UpdateJobApply(jobApply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJobApply 用id查询JobApply
// @Tags JobApply
// @Summary 用id查询JobApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.JobApply true "用id查询JobApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jobApply/findJobApply [get]
func (jobApplyApi *JobApplyApi) FindJobApply(c *gin.Context) {
	var jobApply clothing.JobApply
	err := c.ShouldBindQuery(&jobApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rejobApply, err := jobApplyService.GetJobApply(jobApply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rejobApply": rejobApply}, c)
	}
}

// GetJobApplyList 分页获取JobApply列表
// @Tags JobApply
// @Summary 分页获取JobApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.JobApplySearch true "分页获取JobApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobApply/getJobApplyList [get]
func (jobApplyApi *JobApplyApi) GetJobApplyList(c *gin.Context) {
	var pageInfo clothingReq.JobApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := jobApplyService.GetJobApplyInfoList(pageInfo); err != nil {
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

// OptApply 工单申请审核
// @Tags JobApply
// @Summary 工单申请审核
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.OptJobApply true "工单申请审核"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobApply/optApply [put]
func (jobApplyApi *JobApplyApi) OptApply(c *gin.Context) {
	var opt clothingReq.OptJobApply
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	apply, err := jobApplyService.GetJobApply(uint(opt.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	global.GVA_KEYLOCK.Lock(fmt.Sprintf("cropping-%d", apply.CroppingID))
	defer global.GVA_KEYLOCK.Unlock(fmt.Sprintf("cropping-%d", apply.CroppingID))
	team, err := teamService.GetTeam(apply.TeamID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if team.UserID != userID {
		global.GVA_LOG.Error("权限不足!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := jobApplyService.OptApply(apply, opt.Status); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}
