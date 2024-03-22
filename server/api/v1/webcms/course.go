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
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type CourseApi struct {
}

var courseService = service.ServiceGroupApp.WebcmsServiceGroup.CourseService

// CreateCourse 创建Course
// @Tags Course
// @Summary 创建Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Course true "创建Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /course/createCourse [post]
func (courseApi *CourseApi) CreateCourse(c *gin.Context) {
	var course webcms.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.CreateCourse(course); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCourse 删除Course
// @Tags Course
// @Summary 删除Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Course true "删除Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /course/deleteCourse [delete]
func (courseApi *CourseApi) DeleteCourse(c *gin.Context) {
	var course webcms.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.DeleteCourse(course); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCourseByIds 批量删除Course
// @Tags Course
// @Summary 批量删除Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /course/deleteCourseByIds [delete]
func (courseApi *CourseApi) DeleteCourseByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.DeleteCourseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCourse 更新Course
// @Tags Course
// @Summary 更新Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Course true "更新Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /course/updateCourse [put]
func (courseApi *CourseApi) UpdateCourse(c *gin.Context) {
	var course webcms.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.UpdateCourse(course); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCourse 用id查询Course
// @Tags Course
// @Summary 用id查询Course
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcms.Course true "用id查询Course"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /course/findCourse [get]
func (courseApi *CourseApi) FindCourse(c *gin.Context) {
	var course webcms.Course
	err := c.ShouldBindQuery(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recourse, err := courseService.GetCourse(course.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recourse": recourse}, c)
	}
}

// GetCourseList 分页获取Course列表
// @Tags Course
// @Summary 分页获取Course列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcmsReq.CourseSearch true "分页获取Course列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /course/getCourseList [get]
func (courseApi *CourseApi) GetCourseList(c *gin.Context) {
	var pageInfo webcmsReq.CourseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userinfo := utils.GetUserInfo(c)
	if list, total, err := courseService.GetCourseInfoList(pageInfo, userinfo.AuthorityId, uuid.UUID(userinfo.UUID)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {

		// 遍历
		// userinfo := utils.GetUserInfo(c)
		// temp := make([]webcms.Course, 0)
		// //超级管理员和站点管理员可以查看全部数据
		// exclude := map[uint]bool{
		// 	888:  true,
		// 	9528: true,
		// }
		// if !exclude[userinfo.AuthorityId] {
		// 	for k, v := range list {
		// 		if v.CreatedBy == fmt.Sprint(userinfo.UUID) {
		// 			temp = append(temp, list[k])
		// 		}
		// 	}
		// } else {
		// 	temp = list
		// }
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
