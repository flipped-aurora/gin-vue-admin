import service from '@/utils/request'

// @Summary 创建通知
// @Produce  application/json
// @Param data body {title:"string",content:"string",type:"string",targetType:"string",targetIds:"array"}
// @Router /notice/notification/createNotification [post]
export const createNotification = (data) => {
  return service({
    url: '/notice/notification/createNotification',
    method: 'post',
    data: data
  })
}

// @Summary 获取通知列表
// @Produce  application/json
// @Param data body {page:"int",pageSize:"int",title:"string",type:"string",status:"string"}
// @Router /notice/notification/getNotificationList [post]
export const getNotificationList = (data) => {
  return service({
    url: '/notice/notification/getNotificationList',
    method: 'post',
    data: data
  })
}

// @Summary 获取通知详情
// @Produce  application/json
// @Param id path int true "通知ID"
// @Router /notice/notification/getNotificationById/{id} [get]
export const getNotificationById = (id) => {
  return service({
    url: `/notice/notification/getNotificationById/${id}`,
    method: 'get'
  })
}

// @Summary 更新通知
// @Produce  application/json
// @Param data body {id:"int",title:"string",content:"string",type:"string",targetType:"string",targetIds:"array"}
// @Router /notice/notification/updateNotification [put]
export const updateNotification = (data) => {
  return service({
    url: '/notice/notification/updateNotification',
    method: 'put',
    data: data
  })
}

// @Summary 删除通知
// @Produce  application/json
// @Param id path int true "通知ID"
// @Router /notice/notification/deleteNotification/{id} [delete]
export const deleteNotification = (id) => {
  return service({
    url: `/notice/notification/deleteNotification/${id}`,
    method: 'delete'
  })
}

// @Summary 发送通知
// @Produce  application/json
// @Param data body {id:"int"}
// @Router /notice/notification/sendNotification [post]
export const sendNotification = (data) => {
  return service({
    url: '/notice/notification/sendNotification',
    method: 'post',
    data: data
  })
}

// @Summary 发布通知
// @Produce  application/json
// @Param id path int true "通知ID"
// @Router /notice/notification/publishNotification/{id} [post]
export const publishNotification = (id) => {
  return service({
    url: `/notice/notification/publishNotification/${id}`,
    method: 'post'
  })
}

// @Summary 获取用户通知列表
// @Produce  application/json
// @Param data body {page:"int",pageSize:"int",isRead:"bool"}
// @Router /notice/notification/getUserNotifications [post]
export const getUserNotificationList = (data) => {
  return service({
    url: '/notice/notification/getUserNotifications',
    method: 'post',
    data: data
  })
}

// @Summary 标记通知为已读
// @Produce  application/json
// @Param data body {notificationId:"int"}
// @Router /notice/notification/markAsRead [post]
export const markNotificationAsRead = (data) => {
  return service({
    url: '/notice/notification/markAsRead',
    method: 'post',
    data: data
  })
}

// @Summary 获取未读通知统计
// @Produce  application/json
// @Router /notice/notification/getNotificationStats [get]
export const getUnreadNotificationCount = () => {
  return service({
    url: '/notice/notification/getNotificationStats',
    method: 'get'
  })
}

// @Summary 获取在线用户列表
// @Produce  application/json
// @Param data body {page:"int",pageSize:"int",username:"string",nickName:"string",roleId:"int"}
// @Router /notice/onlineUser/getOnlineUsers [post]
export const getOnlineUserList = (data) => {
  return service({
    url: '/notice/onlineUser/getOnlineUsers',
    method: 'post',
    data: data
  })
}

// @Summary 删除用户通知记录
// @Produce  application/json
// @Param notificationId path int true "通知ID"
// @Router /notice/notification/deleteUserNotification/{notificationId} [delete]
export const deleteUserNotification = (notificationId) => {
  return service({
    url: `/notice/notification/deleteUserNotification/${notificationId}`,
    method: 'delete'
  })
}

// @Summary 强制用户下线
// @Produce  application/json
// @Param userId path int true "用户ID"
// @Router /notice/onlineUser/removeOnlineUser/{userId} [delete]
export const kickUser = (userId) => {
  return service({
    url: `/notice/onlineUser/removeOnlineUser/${userId}`,
    method: 'delete'
  })
}

// @Summary 获取在线用户统计
// @Produce  application/json
// @Router /notice/onlineUser/getOnlineUserStats [get]
export const getOnlineUserStats = () => {
  return service({
    url: '/notice/onlineUser/getOnlineUserStats',
    method: 'get'
  })
}
