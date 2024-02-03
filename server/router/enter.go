package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/comment"
	"github.com/flipped-aurora/gin-vue-admin/server/router/competition"
	"github.com/flipped-aurora/gin-vue-admin/server/router/competitionInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/discuss"
	"github.com/flipped-aurora/gin-vue-admin/server/router/disscussInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
	"github.com/flipped-aurora/gin-vue-admin/server/router/userInfo"
)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	UserInfo        userInfo.RouterGroup
	CompetitionInfo competitionInfo.RouterGroup
	DisscussInfo    disscussInfo.RouterGroup
	User            user.RouterGroup
	Competition     competition.RouterGroup
	Discuss         discuss.RouterGroup
	Comment         comment.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
