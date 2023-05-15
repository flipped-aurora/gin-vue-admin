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

type JobApi struct {
}

var jobService = service.ServiceGroupApp.ClothingServiceGroup.JobService

// CreateJob 创建Job
// @Tags Job
// @Summary 创建Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Job true "创建Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /job/createJob [post]
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

// DeleteJob 删除Job
// @Tags Job
// @Summary 删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Job true "删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /job/deleteJob [delete]
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

// DeleteJobByIds 批量删除Job
// @Tags Job
// @Summary 批量删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /job/deleteJobByIds [delete]
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

// UpdateJob 更新Job
// @Tags Job
// @Summary 更新Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Job true "更新Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /job/updateJob [put]
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

func (jobApi *JobApi) GetJobListByFilter(c *gin.Context) {
	var pageInfo clothingReq.JobSearch
}
