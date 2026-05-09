package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{Path: "/article/createArticle", Description: "创建文章", ApiGroup: "文章管理", Method: "POST"},
		{Path: "/article/deleteArticle", Description: "删除文章", ApiGroup: "文章管理", Method: "DELETE"},
		{Path: "/article/deleteArticleByIds", Description: "批量删除文章", ApiGroup: "文章管理", Method: "DELETE"},
		{Path: "/article/updateArticle", Description: "更新文章", ApiGroup: "文章管理", Method: "PUT"},
		{Path: "/article/findArticle", Description: "用ID查询文章", ApiGroup: "文章管理", Method: "GET"},
		{Path: "/article/readArticle", Description: "阅读文章", ApiGroup: "文章管理", Method: "GET"},
		{Path: "/article/getArticleList", Description: "分页获取文章列表", ApiGroup: "文章管理", Method: "GET"},
		{Path: "/article/createCategory", Description: "创建分类", ApiGroup: "文章分类", Method: "POST"},
		{Path: "/article/deleteCategory", Description: "删除分类", ApiGroup: "文章分类", Method: "DELETE"},
		{Path: "/article/updateCategory", Description: "更新分类", ApiGroup: "文章分类", Method: "PUT"},
		{Path: "/article/getCategoryList", Description: "分页获取分类列表", ApiGroup: "文章分类", Method: "GET"},
		{Path: "/article/getAllCategories", Description: "获取所有分类", ApiGroup: "文章分类", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
