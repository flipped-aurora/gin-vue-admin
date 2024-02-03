package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comment"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/competition"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/discuss"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/user"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	UserApiGroup        user.ApiGroup
	CompetitionApiGroup competition.ApiGroup
	DiscussApiGroup     discuss.ApiGroup
	CommentApiGroup     comment.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
