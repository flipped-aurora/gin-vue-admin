package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/wubi/api"

var (
	Router   = new(router)
	apiScore = api.Api.Score
)

type router struct {
	Score score
}
