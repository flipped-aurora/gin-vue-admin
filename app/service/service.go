package service

import (
	"app/service/comment"
	"app/service/competition"
	"app/service/discuss"
	"app/service/user"
)

type ServiceGroup struct {
	UserService    user.UserInfoService
	ComService     competition.ComInfoService
	DisService     discuss.DisInfoService
	CommentService comment.CommentInfoService
}

var ServiceApi = new(ServiceGroup)
