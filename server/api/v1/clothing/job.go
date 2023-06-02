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

type JobApi struct {
}

var jobService = service.ServiceGroupApp.ClothingServiceGroup.JobService

func (jobApi *JobApi) CreateJob(c *gin.Context) {
	var job clothing.Job
	err := c.ShouldBindJSON(&job)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	job.CreatedBy = utils.GetUserID(c)
	if err := jobService.CreateJob(&job); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (jobApi *JobApi) DeleteJob(c *gin.Context) {
	var job clothing.Job
	err := c.ShouldBindJSON(&job)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	job.DeletedBy = utils.GetUserID(c)
	if err := jobService.DeleteJob(job); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (jobApi *JobApi) DeleteJobByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := jobService.DeleteJobByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (jobApi *JobApi) UpdateJob(c *gin.Context) {
	var job clothing.Job
	err := c.ShouldBindJSON(&job)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	job.UpdatedBy = utils.GetUserID(c)
	if err := jobService.UpdateJob(job); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (jobApi *JobApi) ChangeWorker(c *gin.Context) {
	var job clothing.Job
	err := c.ShouldBindJSON(&job)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	job.UpdatedBy = utils.GetUserID(c)
	if err := jobService.ChangeWorker(job); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJob 用id查询Job
// @Tags Job
// @Summary 用id查询Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Job true "用id查询Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /job/findJob [get]
func (jobApi *JobApi) FindJob(c *gin.Context) {
	var job clothing.Job
	err := c.ShouldBindQuery(&job)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rejob, err := jobService.GetJob(job.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rejob": rejob}, c)
	}
}

// GetJobList 分页获取Job列表
// @Tags Job
// @Summary 分页获取Job列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.JobSearch true "分页获取Job列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /job/getJobList [get]
func (jobApi *JobApi) GetJobList(c *gin.Context) {
	var pageInfo clothingReq.JobSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := jobService.GetJobInfoList(pageInfo); err != nil {
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

func (jobApi *JobApi) PostJobList(c *gin.Context) {
	var data clothingReq.JobList
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	team, err := teamService.GetTeam(data.TeamID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if team.UserID != userID {
		response.FailWithMessage("权限不足", c)
		return
	}
	cropping, err := croppingRecordService.GetCroppingRecord(data.CroppingID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if err := jobService.PostJob(cropping, data); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (jobApi *JobApi) JobAuditApply(c *gin.Context) {
	var job clothingReq.JobAuditApply
	err := c.ShouldBindJSON(&job)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var oJob clothing.Job
	if job.Type == 0 {
		oJob, err = jobService.GetJob(job.ID)
		if err != nil {
			response.FailWithMessage("找不到工单", c)
			return
		}
		if utils.GetUserID(c) != oJob.UserID {
			response.FailWithMessage("权限不足", c)
			return
		}
	} else {
		var process clothing.Process
		if err := global.GVA_DB.First(&process, job.ProcessID).Error; err != nil {
			response.FailWithMessage("工序不存在", c)
			return
		}
		oJob.JobType = enum.Process
		oJob.CroppingID = 0
		oJob.UserID = utils.GetUserID(c)
		oJob.TeamID = job.TeamID
		oJob.Step = enum.CroppingHandling
		oJob.Quantity = int(job.RealQuantity)
		oJob.Income = process.Price * float64(job.RealQuantity)
		oJob.UpdatedBy = utils.GetUserID(c)
		oJob.CreatedBy = utils.GetUserID(c)
		oJob.ProcessID = process.ID
		oJob.ProcessName = process.Name
		oJob.Price = process.Price
		if err := global.GVA_DB.Create(&oJob).Error; err != nil {
			response.FailWithMessage("服务器错误", c)
			return
		}
		oJob, err = jobService.GetJob(oJob.ID)
		if err != nil {
			response.FailWithMessage("找不到工单", c)
			return
		}
	}
	oJob.UpdatedBy = utils.GetUserID(c)
	if err := jobService.AuditApply(oJob, job.RealQuantity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		err := msgBoxService.SendMsg(oJob.UpdatedBy, oJob.Team.UserID, enum.JobComplete, oJob.ID)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
		response.OkWithMessage("更新成功", c)
	}
}

func (jobApi *JobApi) JobAuditOpt(c *gin.Context) {
	var job clothingReq.JobAuditOpt
	err := c.ShouldBindJSON(&job)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	oJob, err := jobService.GetJob(job.ID)
	if err != nil {
		response.FailWithMessage("找不到工单", c)
		return
	}
	oJob.UpdatedBy = utils.GetUserID(c)
	if utils.GetUserID(c) != oJob.Team.UserID {
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := jobService.JobAuditOpt(oJob, job.Status); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
