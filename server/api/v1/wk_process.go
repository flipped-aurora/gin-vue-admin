package v1

import (
	"fmt"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags WorkflowProcess
// @Summary 创建WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "创建WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/createWorkflowProcess [post]
func CreateWorkflowProcess(c *gin.Context) {
	var workflowProcess model.WorkflowProcess
	_ = c.ShouldBindJSON(&workflowProcess)
	err := service.CreateWorkflowProcess(workflowProcess)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags WorkflowProcess
// @Summary 删除WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "删除WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workflowProcess/deleteWorkflowProcess [delete]
func DeleteWorkflowProcess(c *gin.Context) {
	var workflowProcess model.WorkflowProcess
	_ = c.ShouldBindJSON(&workflowProcess)
	err := service.DeleteWorkflowProcess(workflowProcess)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags WorkflowProcess
// @Summary 批量删除WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workflowProcess/deleteWorkflowProcessByIds [delete]
func DeleteWorkflowProcessByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteWorkflowProcessByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags WorkflowProcess
// @Summary 更新WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "更新WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /workflowProcess/updateWorkflowProcess [put]
func UpdateWorkflowProcess(c *gin.Context) {
	var workflowProcess model.WorkflowProcess
	_ = c.ShouldBindJSON(&workflowProcess)
	err := service.UpdateWorkflowProcess(&workflowProcess)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags WorkflowProcess
// @Summary 用id查询WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "用id查询WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workflowProcess/findWorkflowProcess [get]
func FindWorkflowProcess(c *gin.Context) {
	var workflowProcess model.WorkflowProcess
	_ = c.ShouldBindQuery(&workflowProcess)
	err, reworkflowProcess := service.GetWorkflowProcess(workflowProcess.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reworkflowProcess": reworkflowProcess}, c)
	}
}

// @Tags WorkflowProcess
// @Summary 用id查询工作流步骤
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "用id查询WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workflowProcess/findWorkflowCreateStep [get]
func FindWorkflowCreateStep(c *gin.Context) {
	var workflowProcess model.WorkflowProcess
	_ = c.ShouldBindQuery(&workflowProcess)
	err, workflow := service.GetWorkflowCreateStep(workflowProcess.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"workflow": workflow}, c)
	}
}

// @Tags WorkflowProcess
// @Summary 分页获取WorkflowProcess列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.WorkflowProcessSearch true "分页获取WorkflowProcess列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getWorkflowProcessList [get]
func GetWorkflowProcessList(c *gin.Context) {
	var pageInfo request.WorkflowProcessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetWorkflowProcessInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
