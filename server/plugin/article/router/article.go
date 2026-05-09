package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Article = new(article)

type article struct{}

func (r *article) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("article").Use(middleware.OperationRecord())
		group.POST("createArticle", apiArticle.CreateArticle)
		group.DELETE("deleteArticle", apiArticle.DeleteArticle)
		group.DELETE("deleteArticleByIds", apiArticle.DeleteArticleByIds)
		group.PUT("updateArticle", apiArticle.UpdateArticle)
	}
	{
		group := private.Group("article")
		group.GET("findArticle", apiArticle.FindArticle)
		group.GET("readArticle", apiArticle.ReadArticle)
		group.GET("getArticleList", apiArticle.GetArticleList)
	}
}
