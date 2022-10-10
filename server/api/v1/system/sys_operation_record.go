package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// CreateSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   创建SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysOperationRecord      true  "创建SysOperationRecord"
// @Success   200   {object}  response.Response{msg=string}  "创建SysOperationRecord"
// @Router    /sysOperationRecord/createSysOperationRecord [post]
func (s *OperationRecordApi) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindJSON(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.CreateSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.createSuccss"), c)
}

// DeleteSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysOperationRecord      true  "SysOperationRecord模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysOperationRecord"
// @Router    /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindJSON(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deletFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteSysOperationRecordByIds
// @Tags      SysOperationRecord
// @Summary   批量删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "批量删除SysOperationRecord"
// @Success   200   {object}  response.Response{msg=string}  "批量删除SysOperationRecord"
// @Router    /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecordByIds(IDS)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("sys_operation_record.batchDeleteSuccess"), c)
}

// FindSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   用id查询SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysOperationRecord                                  true  "Id"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysOperationRecord"
// @Router    /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindQuery(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(sysOperationRecord, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysOperationRecord, err := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithDetailed(gin.H{"reSysOperationRecord": reSysOperationRecord}, global.Translate("general.querySuccess"), c)
}

// GetSysOperationRecordList
// @Tags      SysOperationRecord
// @Summary   分页获取SysOperationRecord列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SysOperationRecordSearch                        true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := operationRecordService.GetSysOperationRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, global.Translate("general.getDataSuccess"), c)
}
