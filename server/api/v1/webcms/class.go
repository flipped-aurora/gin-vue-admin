package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ClassApi struct {
}

var classService = service.ServiceGroupApp.WebcmsServiceGroup.ClassService

// CreateClass 创建Class
// @Tags Class
// @Summary 创建Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Class true "创建Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/createClass [post]
func (classApi *ClassApi) CreateClass(c *gin.Context) {
	var class webcms.Class
	err := c.ShouldBindJSON(&class)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":     {utils.NotEmpty()},
		"VideoUrl":  {utils.NotEmpty()},
		"TotalHour": {utils.NotEmpty()},
		"TeacherId": {utils.NotEmpty()},
	}
	if err := utils.Verify(class, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMapString("siteinfo")
	class.Siteid = siteinfo["id"]
	if err := classService.CreateClass(class); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClass 删除Class
// @Tags Class
// @Summary 删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Class true "删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /class/deleteClass [delete]
func (classApi *ClassApi) DeleteClass(c *gin.Context) {
	var class webcms.Class
	err := c.ShouldBindJSON(&class)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := classService.DeleteClass(class); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClassByIds 批量删除Class
// @Tags Class
// @Summary 批量删除Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /class/deleteClassByIds [delete]
func (classApi *ClassApi) DeleteClassByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := classService.DeleteClassByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClass 更新Class
// @Tags Class
// @Summary 更新Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Class true "更新Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /class/updateClass [put]
func (classApi *ClassApi) UpdateClass(c *gin.Context) {
	var class webcms.Class
	err := c.ShouldBindJSON(&class)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":     {utils.NotEmpty()},
		"VideoUrl":  {utils.NotEmpty()},
		"TotalHour": {utils.NotEmpty()},
		"TeacherId": {utils.NotEmpty()},
	}
	if err := utils.Verify(class, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := classService.UpdateClass(class); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClass 用id查询Class
// @Tags Class
// @Summary 用id查询Class
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcms.Class true "用id查询Class"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /class/findClass [get]
func (classApi *ClassApi) FindClass(c *gin.Context) {
	var class webcms.Class
	err := c.ShouldBindQuery(&class)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reclass, err := classService.GetClass(class.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclass": reclass}, c)
	}
}

// GetClassList 分页获取Class列表
// @Tags Class
// @Summary 分页获取Class列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcmsReq.ClassSearch true "分页获取Class列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/getClassList [get]
func (classApi *ClassApi) GetClassList(c *gin.Context) {
	var pageInfo webcmsReq.ClassSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMapString("siteinfo")
	if list, total, err := classService.GetClassInfoList(pageInfo, siteinfo["id"]); err != nil {
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

// 通过课程id获取Class列表
func (classApi *ClassApi) GetClassListById(c *gin.Context) {
	course_id := c.DefaultQuery("course_id", "1")
	if list, err := classService.GetClassListById(course_id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
