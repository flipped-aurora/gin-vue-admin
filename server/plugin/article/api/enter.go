package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/article/service"

var (
	Api            = new(api)
	serviceArticle = service.Service.Article
	serviceCategory = service.Service.Category
)

type api struct {
	Article  article
	Category category
}
