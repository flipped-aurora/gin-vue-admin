package NestExecRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg"
	NestExecRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NestExecRecordApi struct {
}

var NtERecordService = service.ServiceGroupApp.NestExecRecordPkgServiceGroup.NestExecRecordService
var NestAirlineService = service.ServiceGroupApp.NestAirlinePkgServiceGroup.NestAirlineService

// CreateNestExecRecord 创建NestExecRecord
// @Tags NestExecRecord
// @Summary 创建NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestExecRecordPkg.NestExecRecord true "创建NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtERecord/createNestExecRecord [post]
func (NtERecordApi *NestExecRecordApi) CreateNestExecRecord(c *gin.Context) {
	var NtERecord NestExecRecordPkg.NestExecRecord
	err := c.ShouldBindJSON(&NtERecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//如果航线名称为空则根据missionid查询航线名称
	if NtERecord.MissionName == "" && NtERecord.Missionid != "" {
		airline, getErr := NestAirlineService.GetNestAirlineByMIssionId(NtERecord.Missionid, c)
		if getErr != nil {
			NtERecord.MissionName = airline.Name
		}
	}

	NtERecord.CreatedBy = utils.GetUserID(c)
	if err := NtERecordService.CreateNestExecRecord(&NtERecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		//因业务需要增加主键id返回
		m := make(map[string]uint)
		m["id"] = NtERecord.ID
		response.OkWithMessageRes("创建成功", m, c)
	}
}

// DeleteNestExecRecord 删除NestExecRecord
// @Tags NestExecRecord
// @Summary 删除NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestExecRecordPkg.NestExecRecord true "删除NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NtERecord/deleteNestExecRecord [delete]
func (NtERecordApi *NestExecRecordApi) DeleteNestExecRecord(c *gin.Context) {
	var NtERecord NestExecRecordPkg.NestExecRecord
	err := c.ShouldBindJSON(&NtERecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	NtERecord.DeletedBy = utils.GetUserID(c)
	if err := NtERecordService.DeleteNestExecRecord(NtERecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNestExecRecordByIds 批量删除NestExecRecord
// @Tags NestExecRecord
// @Summary 批量删除NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /NtERecord/deleteNestExecRecordByIds [delete]
func (NtERecordApi *NestExecRecordApi) DeleteNestExecRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := NtERecordService.DeleteNestExecRecordByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNestExecRecord 更新NestExecRecord
// @Tags NestExecRecord
// @Summary 更新NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestExecRecordPkg.NestExecRecord true "更新NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NtERecord/updateNestExecRecord [put]
func (NtERecordApi *NestExecRecordApi) UpdateNestExecRecord(c *gin.Context) {
	var NtERecord NestExecRecordPkg.NestExecRecord
	err := c.ShouldBindJSON(&NtERecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	NtERecord.UpdatedBy = utils.GetUserID(c)
	if err := NtERecordService.UpdateNestExecRecord(NtERecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateNestExecRecord 更新NestExecRecord及关联的航线数据
// @Tags NestExecRecord
// @Summary 更新NestExecRecord及关联的航线数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestExecRecordPkg.NestExecRecord true "更新NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NtERecord/updateNestExecRecord [put]
func (NtERecordApi *NestExecRecordApi) UpdateNestExecRecordWithAirline(c *gin.Context) {
	var NtERecord NestExecRecordPkg.NestExecRecord
	err := c.ShouldBindJSON(&NtERecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if NtERecord.Missionid != "" {
		airlineRes, getErr := NestAirlineService.GetNestAirlineByMIssionId(NtERecord.Missionid, c)
		if getErr != nil {
			//todo add nest_exec_record query airline msg exception
		} else {
			NtERecord.MissionName = airlineRes.Name
		}
	}
	NtERecord.UpdatedBy = utils.GetUserID(c)
	if err := NtERecordService.UpdateNestExecRecord(NtERecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNestExecRecord 用id查询NestExecRecord
// @Tags NestExecRecord
// @Summary 用id查询NestExecRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NestExecRecordPkg.NestExecRecord true "用id查询NestExecRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NtERecord/findNestExecRecord [get]
func (NtERecordApi *NestExecRecordApi) FindNestExecRecord(c *gin.Context) {
	var NtERecord NestExecRecordPkg.NestExecRecord
	err := c.ShouldBindQuery(&NtERecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reNtERecord, err := NtERecordService.GetNestExecRecord(NtERecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reNtERecord": reNtERecord}, c)
	}
}

// GetNestExecRecordList 分页获取NestExecRecord列表
// @Tags NestExecRecord
// @Summary 分页获取NestExecRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NestExecRecordPkgReq.NestExecRecordSearch true "分页获取NestExecRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtERecord/getNestExecRecordList [get]
func (NtERecordApi *NestExecRecordApi) GetNestExecRecordList(c *gin.Context) {
	var pageInfo NestExecRecordPkgReq.NestExecRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := NtERecordService.GetNestExecRecordInfoList(pageInfo, c); err != nil {
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
