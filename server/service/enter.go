package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/comment"
	"github.com/flipped-aurora/gin-vue-admin/server/service/competition"
	"github.com/flipped-aurora/gin-vue-admin/server/service/discuss"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	UserServiceGroup        user.ServiceGroup
	CompetitionServiceGroup competition.ServiceGroup
	DiscussServiceGroup     discuss.ServiceGroup
	CommentServiceGroup     comment.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
