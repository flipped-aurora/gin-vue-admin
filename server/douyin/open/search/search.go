package search

import "github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"

// Search 搜索管理
type Search struct {
	*context.Context
}

// NewSearch .
func NewSearch(context *context.Context) *Search {
	search := new(Search)
	search.Context = context
	return search
}
