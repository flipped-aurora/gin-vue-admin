package poi

import "github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"

// Poi 生活服务开放能力
type Poi struct {
	*context.Context
}

// NewPoi .
func NewPoi(context *context.Context) *Poi {
	poi := new(Poi)
	poi.Context = context
	return poi
}
