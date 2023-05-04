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

type JobQuestionApi struct {
}

var jobQuestionService = service.ServiceGroupApp.ClothingServiceGroup.JobQuestionService


// CreateJobQuestion 创建JobQuestion
// @Tags JobQuestion
// @Summary 创建JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.JobQuestion true "创建JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobQuestion/createJobQuestion [post]
func (jobQuestionApi *JobQuestionApi) CreateJobQuestion(c *gin.Context) {
	var jobQuestion clothing.JobQuestion
	err := c.ShouldBindJSON(&jobQuestion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    jobQuestion.CreatedBy = utils.GetUserID(c)
	if err := jobQuestionService.CreateJobQuestion(&jobQuestion); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJobQuestion 删除JobQuestion
// @Tags JobQuestion
// @Summary 删除JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.JobQuestion true "删除JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jobQuestion/deleteJobQuestion [delete]
func (jobQuestionApi *JobQuestionApi) DeleteJobQuestion(c *gin.Context) {
	var jobQuestion clothing.JobQuestion
	err := c.ShouldBindJSON(&jobQuestion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    jobQuestion.DeletedBy = utils.GetUserID(c)
	if err := jobQuestionService.DeleteJobQuestion(jobQuestion); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJobQuestionByIds 批量删除JobQuestion
// @Tags JobQuestion
// @Summary 批量删除JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /jobQuestion/deleteJobQuestionByIds [delete]
func (jobQuestionApi *JobQuestionApi) DeleteJobQuestionByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := jobQuestionService.DeleteJobQuestionByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJobQuestion 更新JobQuestion
// @Tags JobQuestion
// @Summary 更新JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.JobQuestion true "更新JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jobQuestion/updateJobQuestion [put]
func (jobQuestionApi *JobQuestionApi) UpdateJobQuestion(c *gin.Context) {
	var jobQuestion clothing.JobQuestion
	err := c.ShouldBindJSON(&jobQuestion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    jobQuestion.UpdatedBy = utils.GetUserID(c)
	if err := jobQuestionService.UpdateJobQuestion(jobQuestion); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJobQuestion 用id查询JobQuestion
// @Tags JobQuestion
// @Summary 用id查询JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.JobQuestion true "用id查询JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jobQuestion/findJobQuestion [get]
func (jobQuestionApi *JobQuestionApi) FindJobQuestion(c *gin.Context) {
	var jobQuestion clothing.JobQuestion
	err := c.ShouldBindQuery(&jobQuestion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rejobQuestion, err := jobQuestionService.GetJobQuestion(jobQuestion.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rejobQuestion": rejobQuestion}, c)
	}
}

// GetJobQuestionList 分页获取JobQuestion列表
// @Tags JobQuestion
// @Summary 分页获取JobQuestion列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.JobQuestionSearch true "分页获取JobQuestion列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobQuestion/getJobQuestionList [get]
func (jobQuestionApi *JobQuestionApi) GetJobQuestionList(c *gin.Context) {
	var pageInfo clothingReq.JobQuestionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := jobQuestionService.GetJobQuestionInfoList(pageInfo); err != nil {
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
