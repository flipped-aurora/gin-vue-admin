package interactive

import "github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"

// Interactive 互动管理
type Interactive struct {
	*context.Context
}

// NewInteractive .
func NewInteractive(context *context.Context) *Interactive {
	interactive := new(Interactive)
	interactive.Context = context
	return interactive
}
