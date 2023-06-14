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

type ProcessApi struct {
}

var processService = service.ServiceGroupApp.ClothingServiceGroup.ProcessService

// CreateProcess 创建Process
// @Tags Process
// @Summary 创建Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Process true "创建Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /process/createProcess [post]
func (processApi *ProcessApi) CreateProcess(c *gin.Context) {
	var process clothingReq.CreateProcess
	err := c.ShouldBindJSON(&process)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	style, err := styleService.GetStyle(process.StyleID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("款式不存在", c)
		return
	}
	if !userRoleService.CheckManager(utils.GetUserID(c), style.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := processService.CreateProcessBatch(process); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProcess 删除Process
// @Tags Process
// @Summary 删除Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Process true "删除Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /process/deleteProcess [delete]
func (processApi *ProcessApi) DeleteProcess(c *gin.Context) {
	var process clothing.Process
	err := c.ShouldBindJSON(&process)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	process.DeletedBy = utils.GetUserID(c)
	style, err := styleService.GetStyle(process.StyleID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("款式不存在", c)
		return
	}
	if !userRoleService.CheckManager(utils.GetUserID(c), style.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := processService.DeleteProcess(process); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (processApi *ProcessApi) DeleteProcessByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := processService.DeleteProcessByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProcess 更新Process
// @Tags Process
// @Summary 更新Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Process true "更新Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /process/updateProcess [put]
func (processApi *ProcessApi) UpdateProcess(c *gin.Context) {
	var process clothing.Process
	err := c.ShouldBindJSON(&process)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	process.UpdatedBy = utils.GetUserID(c)
	style, err := styleService.GetStyle(process.StyleID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("款式不存在", c)
		return
	}
	if !userRoleService.CheckManager(utils.GetUserID(c), style.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := processService.UpdateProcess(process); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProcess 用id查询Process
// @Tags Process
// @Summary 用id查询Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Process true "用id查询Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /process/findProcess [get]
func (processApi *ProcessApi) FindProcess(c *gin.Context) {
	var process clothing.Process
	err := c.ShouldBindQuery(&process)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reprocess, err := processService.GetProcess(process.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprocess": reprocess}, c)
	}
}

// GetProcessList 分页获取Process列表
// @Tags Process
// @Summary 分页获取Process列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.ProcessSearch true "分页获取Process列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /process/getProcessList [get]
func (processApi *ProcessApi) GetProcessList(c *gin.Context) {
	var pageInfo clothingReq.ProcessSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := processService.GetProcessInfoList(pageInfo); err != nil {
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
