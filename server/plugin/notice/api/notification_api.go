package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/request"
	noticeService "github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type NotificationApi struct{}

var notificationService = noticeService.ServiceGroupApp.NotificationService

// CreateNotification 创建通知
// @Tags     NoticeCenter
// @Summary  创建一个新的通知
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.CreateNotificationRequest true "通知的标题、内容、类型等信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /notice/createNotification [post]
func (a *NotificationApi) CreateNotification(c *gin.Context) {
	var req request.CreateNotificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	userId := utils.GetUserID(c)

	notification, err := notificationService.CreateNotification(req, userId)
	if err != nil {
		global.GVA_LOG.Error("创建通知失败!" + err.Error())
		response.FailWithMessage("创建通知失败", c)
		return
	}

	response.OkWithDetailed(notification, "创建成功", c)
}

// UpdateNotification 更新通知
// @Tags     NoticeCenter
// @Summary  更新通知信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.UpdateNotificationRequest true "通知的更新信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /notice/updateNotification [put]
func (a *NotificationApi) UpdateNotification(c *gin.Context) {
	var req request.UpdateNotificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = notificationService.UpdateNotification(req)
	if err != nil {
		global.GVA_LOG.Error("更新通知失败!" + err.Error())
		response.FailWithMessage("更新通知失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteNotification 删除通知
// @Tags     NoticeCenter
// @Summary  删除通知
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "通知ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /notice/deleteNotification/{id} [delete]
func (a *NotificationApi) DeleteNotification(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的通知ID", c)
		return
	}

	err = notificationService.DeleteNotification(uint(id))
	if err != nil {
		global.GVA_LOG.Error("删除通知失败!" + err.Error())
		response.FailWithMessage("删除通知失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetNotificationList 获取通知列表
// @Tags     NoticeCenter
// @Summary  获取通知列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.NotificationSearch true "分页和搜索条件"
// @Success  200  {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /notice/getNotificationList [post]
func (a *NotificationApi) GetNotificationList(c *gin.Context) {
	var req request.NotificationSearch
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := notificationService.GetNotificationList(req)
	if err != nil {
		global.GVA_LOG.Error("获取通知列表失败!" + err.Error())
		response.FailWithMessage("获取通知列表失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)
}

// GetNotificationById 根据ID获取通知
// @Tags     NoticeCenter
// @Summary  根据ID获取通知详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "通知ID"
// @Success  200  {object} response.Response{data=model.Notification,msg=string} "获取成功"
// @Router   /notice/getNotification/{id} [get]
func (a *NotificationApi) GetNotificationById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的通知ID", c)
		return
	}

	notification, err := notificationService.GetNotificationById(uint(id))
	if err != nil {
		global.GVA_LOG.Error("获取通知失败!" + err.Error())
		response.FailWithMessage("获取通知失败", c)
		return
	}

	response.OkWithDetailed(notification, "获取成功", c)
}

// PublishNotification 发布通知
// @Tags     NoticeCenter
// @Summary  发布通知
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "通知ID"
// @Success  200  {object} response.Response{msg=string} "发布成功"
// @Router   /notice/publishNotification/{id} [post]
func (a *NotificationApi) PublishNotification(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的通知ID", c)
		return
	}

	err = notificationService.PublishNotification(uint(id))
	if err != nil {
		global.GVA_LOG.Error("发布通知失败!" + err.Error())
		response.FailWithMessage("发布通知失败", c)
		return
	}

	response.OkWithMessage("发布成功", c)
}

// SendNotification 发送通知
// @Tags     NoticeCenter
// @Summary  发送通知给指定用户或角色
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.SendNotificationRequest true "发送通知的请求参数"
// @Success  200  {object} response.Response{data=response.SendNotificationResponse,msg=string} "发送成功"
// @Router   /notice/sendNotification [post]
func (a *NotificationApi) SendNotification(c *gin.Context) {
	var req request.SendNotificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	userId := utils.GetUserID(c)

	result, err := notificationService.SendNotification(req, userId)
	if err != nil {
		global.GVA_LOG.Error("发送通知失败!" + err.Error())
		response.FailWithMessage("发送通知失败", c)
		return
	}

	response.OkWithDetailed(result, "发送成功", c)
}

// GetUserNotifications 获取用户通知列表
// @Tags     NoticeCenter
// @Summary  获取当前用户的通知列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.UserNotificationSearch true "分页和搜索条件"
// @Success  200  {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /notice/getUserNotifications [post]
func (a *NotificationApi) GetUserNotifications(c *gin.Context) {
	var req request.UserNotificationSearch
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	userId := utils.GetUserID(c)

	list, total, err := notificationService.GetUserNotifications(userId, req)
	if err != nil {
		global.GVA_LOG.Error("获取用户通知列表失败!" + err.Error())
		response.FailWithMessage("获取用户通知列表失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)
}

// MarkNotificationAsRead 标记通知为已读
// @Tags     NoticeCenter
// @Summary  标记通知为已读
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.MarkReadRequest true "标记已读的请求参数"
// @Success  200  {object} response.Response{msg=string} "标记成功"
// @Router   /notice/markAsRead [post]
func (a *NotificationApi) MarkNotificationAsRead(c *gin.Context) {
	var req request.MarkReadRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	userId := utils.GetUserID(c)

	err = notificationService.MarkNotificationsAsRead(userId, req.NotificationIds)
	if err != nil {
		global.GVA_LOG.Error("标记通知已读失败!" + err.Error())
		response.FailWithMessage("标记通知已读失败", c)
		return
	}

	response.OkWithMessage("标记成功", c)
}

// GetNotificationStats 获取通知统计信息
// @Tags     NoticeCenter
// @Summary  获取通知统计信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=response.NotificationStatsResponse,msg=string} "获取成功"
// @Router   /notice/getNotificationStats [get]
func (a *NotificationApi) GetNotificationStats(c *gin.Context) {
	// 获取当前用户ID
	userId := utils.GetUserID(c)

	stats, err := notificationService.GetNotificationStats(userId)
	if err != nil {
		global.GVA_LOG.Error("获取通知统计失败!" + err.Error())
		response.FailWithMessage("获取通知统计失败", c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// DeleteUserNotification 删除用户通知记录
// @Tags     NoticeCenter
// @Summary  删除用户通知记录，如果通知只针对单个用户则删除通知本身
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    notificationId path int true "通知ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /notice/deleteUserNotification/{notificationId} [delete]
func (a *NotificationApi) DeleteUserNotification(c *gin.Context) {
	notificationIdStr := c.Param("notificationId")
	notificationId, err := strconv.ParseUint(notificationIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("通知ID格式错误", c)
		return
	}

	// 获取当前用户ID
	userId := utils.GetUserID(c)

	err = notificationService.DeleteUserNotification(userId, uint(notificationId))
	if err != nil {
		global.GVA_LOG.Error("删除用户通知失败!" + err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
