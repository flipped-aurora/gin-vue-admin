package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/article/api"

var (
	Router         = new(router)
	apiArticle     = api.Api.Article
	apiCategory    = api.Api.Category
)

type router struct {
	Article  article
	Category category
}
