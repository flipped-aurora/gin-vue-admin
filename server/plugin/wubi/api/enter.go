package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/wubi/service"

var (
	Api          = new(api)
	serviceScore = service.Service.Score
)

type api struct {
	Score score
}
