package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type SysPositionApi struct{}

// CreateSysPosition
// @Tags      SysPosition
// @Summary   创建岗位
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysPosition             true  "岗位名称, 岗位编码"
// @Success   200   {object}  response.Response{msg=string}  "创建岗位"
// @Router    /position/createPosition [post]
func (a *SysPositionApi) CreateSysPosition(c *gin.Context) {
	var pos system.SysPosition
	if err := c.ShouldBindJSON(&pos); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pos.Name == "" {
		response.FailWithMessage("岗位名称不能为空", c)
		return
	}
	if err := positionService.CreateSysPosition(c.Request.Context(), &pos); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("创建失败!")
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateSysPosition
// @Tags      SysPosition
// @Summary   更新岗位
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysPosition             true  "岗位ID, 岗位名称"
// @Success   200   {object}  response.Response{msg=string}  "更新岗位"
// @Router    /position/updatePosition [put]
func (a *SysPositionApi) UpdateSysPosition(c *gin.Context) {
	var pos system.SysPosition
	if err := c.ShouldBindJSON(&pos); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pos.ID == 0 {
		response.FailWithMessage("岗位ID不能为空", c)
		return
	}
	if err := positionService.UpdateSysPosition(c.Request.Context(), &pos); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("更新失败!")
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteSysPosition
// @Tags      SysPosition
// @Summary   删除岗位
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "岗位ID"
// @Success   200   {object}  response.Response{msg=string}  "删除岗位"
// @Router    /position/deletePosition [delete]
func (a *SysPositionApi) DeleteSysPosition(c *gin.Context) {
	var req request.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := positionService.DeleteSysPosition(c.Request.Context(), req.Uint()); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("删除失败!")
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetSysPositionList
// @Tags      SysPosition
// @Summary   分页获取岗位列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SysPositionSearch                                          true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]system.SysPosition},msg=string}  "分页获取岗位列表"
// @Router    /position/getPositionList [post]
func (a *SysPositionApi) GetSysPositionList(c *gin.Context) {
	var search systemReq.SysPositionSearch
	if err := c.ShouldBindJSON(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := positionService.GetSysPositionList(c.Request.Context(), search)
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

// FindSysPosition
// @Tags      SysPosition
// @Summary   根据ID获取岗位
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  query     uint                                                   true  "岗位ID"
// @Success   200  {object}  response.Response{data=system.SysPosition,msg=string}  "获取成功"
// @Router    /position/findPosition [get]
func (a *SysPositionApi) FindSysPosition(c *gin.Context) {
	var req request.GetById
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pos, err := positionService.GetSysPosition(c.Request.Context(), req.Uint())
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(pos, "获取成功", c)
}

// GetPositionUsers
// @Tags      SysPosition
// @Summary   获取岗位下的用户ID列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     positionId  query     uint                                       true  "岗位ID"
// @Success   200         {object}  response.Response{data=[]uint,msg=string}  "获取成功"
// @Router    /position/getPositionUsers [get]
func (a *SysPositionApi) GetPositionUsers(c *gin.Context) {
	var req systemReq.SetPositionUsers
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ids, err := positionService.GetPositionUserIds(c.Request.Context(), req.PositionId)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	if ids == nil {
		ids = []uint{}
	}
	response.OkWithDetailed(ids, "获取成功", c)
}

// SetPositionUsers
// @Tags      SysPosition
// @Summary   设置岗位成员(反向分配用户)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetPositionUsers     true  "岗位ID, 用户ID列表"
// @Success   200   {object}  response.Response{msg=string}  "设置成功"
// @Router    /position/setPositionUsers [post]
func (a *SysPositionApi) SetPositionUsers(c *gin.Context) {
	var req systemReq.SetPositionUsers
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.PositionId == 0 {
		response.FailWithMessage("岗位ID不能为空", c)
		return
	}
	if err := positionService.SetPositionUsers(c.Request.Context(), req.PositionId, req.UserIds); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("设置失败!")
		response.FailWithMessage("设置失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
