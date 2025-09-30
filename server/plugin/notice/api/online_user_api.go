package api

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/request"
	noticeResponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/response"
	noticeService "github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/service"
	"github.com/gin-gonic/gin"
)

type OnlineUserApi struct{}

var onlineUserService = noticeService.ServiceGroupApp.OnlineUserService

// GetOnlineUsers 获取在线用户列表
// @Tags     NoticeCenter
// @Summary  获取在线用户列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]response.OnlineUserResponse,msg=string} "获取成功"
// @Router   /notice/onlineUser/getOnlineUsers [get]
func (a *OnlineUserApi) GetOnlineUsers(c *gin.Context) {
	onlineUsers, err := onlineUserService.GetOnlineUsers()
	if err != nil {
		global.GVA_LOG.Error("获取在线用户列表失败!" + err.Error())
		response.FailWithMessage("获取在线用户列表失败", c)
		return
	}

	response.OkWithDetailed(onlineUsers, "获取成功", c)
}

// GetOnlineUsersWithPagination 获取在线用户列表（分页）
// @Tags     NoticeCenter
// @Summary  获取在线用户列表（支持分页和搜索）
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.OnlineUserSearch true "分页参数"
// @Success  200  {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /notice/onlineUser/getOnlineUsers [post]
func (a *OnlineUserApi) GetOnlineUsersWithPagination(c *gin.Context) {
	var req request.OnlineUserSearch
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	onlineUsers, total, err := onlineUserService.GetOnlineUsersWithPagination(req)
	if err != nil {
		global.GVA_LOG.Error("获取在线用户列表失败!" + err.Error())
		response.FailWithMessage("获取在线用户列表失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     onlineUsers,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)
}

// GetOnlineUsersByRole 根据角色获取在线用户
// @Tags     NoticeCenter
// @Summary  根据角色获取在线用户列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    roleId path int true "角色ID"
// @Success  200  {object} response.Response{data=[]response.OnlineUserResponse,msg=string} "获取成功"
// @Router   /notice/getOnlineUsersByRole/{roleId} [get]
func (a *OnlineUserApi) GetOnlineUsersByRole(c *gin.Context) {
	roleIdStr := c.Param("roleId")
	roleId, err := strconv.ParseUint(roleIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的角色ID", c)
		return
	}

	// 根据角色获取用户ID列表
	userIds, err := onlineUserService.GetUsersByRoleIds([]uint{uint(roleId)})
	if err != nil {
		global.GVA_LOG.Error("获取角色用户失败!" + err.Error())
		response.FailWithMessage("获取角色用户失败", c)
		return
	}

	// 获取所有在线用户
	allOnlineUsers, err := onlineUserService.GetOnlineUsers()
	if err != nil {
		global.GVA_LOG.Error("获取在线用户失败!" + err.Error())
		response.FailWithMessage("获取在线用户失败", c)
		return
	}

	// 过滤出指定角色的在线用户
	var onlineUsers []noticeResponse.OnlineUserResponse
	userIdMap := make(map[uint]bool)
	for _, id := range userIds {
		userIdMap[id] = true
	}
	for _, user := range allOnlineUsers {
		if userIdMap[user.UserId] {
			onlineUsers = append(onlineUsers, user)
		}
	}
	if err != nil {
		global.GVA_LOG.Error("根据角色获取在线用户失败!" + err.Error())
		response.FailWithMessage("根据角色获取在线用户失败", c)
		return
	}

	response.OkWithDetailed(onlineUsers, "获取成功", c)
}

// CheckUserOnline 检查用户是否在线
// @Tags     NoticeCenter
// @Summary  检查指定用户是否在线
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId path int true "用户ID"
// @Success  200  {object} response.Response{data=bool,msg=string} "检查成功"
// @Router   /notice/checkUserOnline/{userId} [get]
func (a *OnlineUserApi) CheckUserOnline(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的用户ID", c)
		return
	}

	isOnline := onlineUserService.IsUserOnline(uint(userId))
	response.OkWithDetailed(isOnline, "检查成功", c)
}

// GetOnlineUserCount 获取在线用户数量
// @Tags     NoticeCenter
// @Summary  获取当前在线用户数量
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=int64,msg=string} "获取成功"
// @Router   /notice/getOnlineUserCount [get]
func (a *OnlineUserApi) GetOnlineUserCount(c *gin.Context) {
	count, err := onlineUserService.GetOnlineUserCount()
	if err != nil {
		global.GVA_LOG.Error("获取在线用户数量失败!" + err.Error())
		response.FailWithMessage("获取在线用户数量失败", c)
		return
	}

	response.OkWithDetailed(count, "获取成功", c)
}

// GetOnlineUserStats 获取在线用户统计数据
// @Tags     NoticeCenter
// @Summary  获取在线用户统计数据（总在线数、今日登录、峰值在线、平均在线）
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=response.OnlineUserStatsResponse,msg=string} "获取成功"
// @Router   /notice/onlineUser/getOnlineUserStats [get]
func (a *OnlineUserApi) GetOnlineUserStats(c *gin.Context) {
	stats, err := onlineUserService.GetOnlineUserStats()
	if err != nil {
		global.GVA_LOG.Error("获取在线用户统计数据失败!" + err.Error())
		response.FailWithMessage("获取在线用户统计数据失败", c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// RemoveOnlineUser 移除在线用户
// @Tags     NoticeCenter
// @Summary  移除指定的在线用户
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId path int true "用户ID"
// @Success  200  {object} response.Response{msg=string} "移除成功"
// @Router   /notice/removeOnlineUser/{userId} [delete]
func (a *OnlineUserApi) RemoveOnlineUser(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的用户ID", c)
		return
	}

	//// 检查权限：只有管理员或用户本人可以移除
	//currentUserId := utils.GetUserID(c)
	//if currentUserId != uint(userId) {
	//	// 这里可以添加管理员权限检查
	//	// 暂时允许所有用户操作，实际项目中应该添加权限验证
	//}

	err = onlineUserService.RemoveOnlineUser(uint(userId))
	if err != nil {
		global.GVA_LOG.Error("移除在线用户失败!" + err.Error())
		response.FailWithMessage("移除在线用户失败", c)
		return
	}

	response.OkWithMessage("移除成功", c)
}

// CleanOfflineUsers 清理离线用户
// @Tags     NoticeCenter
// @Summary  清理长时间未活跃的离线用户
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{msg=string} "清理成功"
// @Router   /notice/cleanOfflineUsers [post]
func (a *OnlineUserApi) CleanOfflineUsers(c *gin.Context) {
	// 清理30分钟未活跃的用户
	err := onlineUserService.CleanOfflineUsers(30 * time.Minute)
	if err != nil {
		global.GVA_LOG.Error("清理离线用户失败!" + err.Error())
		response.FailWithMessage("清理离线用户失败", c)
		return
	}

	response.OkWithMessage("清理成功", c)
}
