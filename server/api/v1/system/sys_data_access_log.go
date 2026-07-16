package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type SysDataAccessLogApi struct{}

// GetDataAccessLogList
// @Tags      SysDataAccessLog
// @Summary   分页获取数据权限审计日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SysDataAccessLogSearch                                                        true  "页码, 每页大小, 事件类型, 表名"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]system.SysDataAccessLog},msg=string}  "分页获取数据权限审计日志"
// @Router    /dataAccessLog/getDataAccessLogList [post]
func (a *SysDataAccessLogApi) GetDataAccessLogList(c *gin.Context) {
	var search systemReq.SysDataAccessLogSearch
	if err := c.ShouldBindJSON(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dataAccessLogService.GetDataAccessLogList(c.Request.Context(), search)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     search.Page,
		PageSize: search.PageSize,
	}, "获取成功", c)
}

// DeleteDataAccessLogByIds
// @Tags      SysDataAccessLog
// @Summary   批量删除数据权限审计日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID列表"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /dataAccessLog/deleteDataAccessLogByIds [delete]
func (a *SysDataAccessLogApi) DeleteDataAccessLogByIds(c *gin.Context) {
	var req request.IdsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dataAccessLogService.DeleteDataAccessLogByIds(c.Request.Context(), req.Ids); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("删除失败!")
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
