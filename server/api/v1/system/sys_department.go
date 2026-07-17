package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type SysDepartmentApi struct{}

// CreateSysDepartment
// @Tags      SysDepartment
// @Summary   创建部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDepartment           true  "部门名称, 父部门ID"
// @Success   200   {object}  response.Response{msg=string}  "创建部门"
// @Router    /department/createDepartment [post]
func (a *SysDepartmentApi) CreateSysDepartment(c *gin.Context) {
	var dept system.SysDepartment
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if dept.Name == "" {
		response.FailWithMessage("部门名称不能为空", c)
		return
	}
	if err := departmentService.CreateSysDepartment(c.Request.Context(), &dept); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("创建失败!")
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateSysDepartment
// @Tags      SysDepartment
// @Summary   更新部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDepartment           true  "部门ID, 部门名称, 父部门ID"
// @Success   200   {object}  response.Response{msg=string}  "更新部门"
// @Router    /department/updateDepartment [put]
func (a *SysDepartmentApi) UpdateSysDepartment(c *gin.Context) {
	var dept system.SysDepartment
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if dept.ID == 0 {
		response.FailWithMessage("部门ID不能为空", c)
		return
	}
	if err := departmentService.UpdateSysDepartment(c.Request.Context(), &dept); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("更新失败!")
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteSysDepartment
// @Tags      SysDepartment
// @Summary   删除部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "部门ID"
// @Success   200   {object}  response.Response{msg=string}  "删除部门"
// @Router    /department/deleteDepartment [delete]
func (a *SysDepartmentApi) DeleteSysDepartment(c *gin.Context) {
	var req request.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := departmentService.DeleteSysDepartment(c.Request.Context(), req.Uint()); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("删除失败!")
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetSysDepartmentList
// @Tags      SysDepartment
// @Summary   获取部门树
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SysDepartmentSearch                       true  "部门名称"
// @Success   200   {object}  response.Response{data=[]system.SysDepartment,msg=string}  "获取部门树"
// @Router    /department/getDepartmentList [post]
func (a *SysDepartmentApi) GetSysDepartmentList(c *gin.Context) {
	var search systemReq.SysDepartmentSearch
	_ = c.ShouldBindJSON(&search)
	list, err := departmentService.GetSysDepartmentTree(c.Request.Context(), search.Name)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}

// FindSysDepartment
// @Tags      SysDepartment
// @Summary   根据ID获取部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  query     uint                                                     true  "部门ID"
// @Success   200  {object}  response.Response{data=system.SysDepartment,msg=string}  "获取成功"
// @Router    /department/findDepartment [get]
func (a *SysDepartmentApi) FindSysDepartment(c *gin.Context) {
	var req request.GetById
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dept, err := departmentService.GetSysDepartment(c.Request.Context(), req.Uint())
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(dept, "获取成功", c)
}

// GetDepartmentUsers
// @Tags      SysDepartment
// @Summary   获取部门下的用户ID列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     deptId  query     uint                                       true  "部门ID"
// @Success   200     {object}  response.Response{data=[]uint,msg=string}  "获取成功"
// @Router    /department/getDepartmentUsers [get]
func (a *SysDepartmentApi) GetDepartmentUsers(c *gin.Context) {
	var req systemReq.SetDepartmentUsers
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ids, err := departmentService.GetDepartmentUserIds(c.Request.Context(), req.DeptId)
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

// SetDepartmentUsers
// @Tags      SysDepartment
// @Summary   设置部门成员(反向分配用户)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetDepartmentUsers   true  "部门ID, 用户ID列表"
// @Success   200   {object}  response.Response{msg=string}  "设置成功"
// @Router    /department/setDepartmentUsers [post]
func (a *SysDepartmentApi) SetDepartmentUsers(c *gin.Context) {
	var req systemReq.SetDepartmentUsers
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.DeptId == 0 {
		response.FailWithMessage("部门ID不能为空", c)
		return
	}
	if err := departmentService.SetDepartmentUsers(c.Request.Context(), req.DeptId, req.UserIds); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("设置失败!")
		response.FailWithMessage("设置失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
