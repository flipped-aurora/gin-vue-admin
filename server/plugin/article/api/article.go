package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Article = new(article)

type article struct{}

// CreateArticle 创建文章
// @Tags     Article
// @Summary  创建文章
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body model.Article true "文章信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /article/createArticle [post]
func (a *article) CreateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	article.UserID = utils.GetUserID(c)
	article.NickName = utils.GetUserInfo(c).NickName
	if err := serviceArticle.CreateArticle(&article); err != nil {
		global.GVA_LOG.Error("创建文章失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteArticle 删除文章
// @Tags     Article
// @Summary  删除文章
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    ID query string true "文章ID"
// @Success  200 {object} response.Response{msg=string} "删除成功"
// @Router   /article/deleteArticle [delete]
func (a *article) DeleteArticle(c *gin.Context) {
	id := c.Query("ID")
	if id == "" {
		response.FailWithMessage("ID必填", c)
		return
	}
	if err := serviceArticle.DeleteArticle(id); err != nil {
		global.GVA_LOG.Error("删除文章失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteArticleByIds 批量删除文章
// @Tags     Article
// @Summary  批量删除文章
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    IDs query []string true "文章ID列表"
// @Success  200 {object} response.Response{msg=string} "批量删除成功"
// @Router   /article/deleteArticleByIds [delete]
func (a *article) DeleteArticleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := serviceArticle.DeleteArticleByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateArticle 更新文章
// @Tags     Article
// @Summary  更新文章
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body model.Article true "文章信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /article/updateArticle [put]
func (a *article) UpdateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serviceArticle.UpdateArticle(article); err != nil {
		global.GVA_LOG.Error("更新文章失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindArticle 用id查询文章(后台管理用,不增加阅读数)
// @Tags     Article
// @Summary  用id查询文章
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    ID query string true "文章ID"
// @Success  200 {object} response.Response{data=model.Article,msg=string} "查询成功"
// @Router   /article/findArticle [get]
func (a *article) FindArticle(c *gin.Context) {
	id := c.Query("ID")
	if id == "" {
		response.FailWithMessage("ID必填", c)
		return
	}
	article, err := serviceArticle.GetArticle(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(article, c)
}

// ReadArticle 阅读文章(增加阅读数)
// @Tags     Article
// @Summary  阅读文章
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    ID query string true "文章ID"
// @Success  200 {object} response.Response{data=model.Article,msg=string} "查询成功"
// @Router   /article/readArticle [get]
func (a *article) ReadArticle(c *gin.Context) {
	id := c.Query("ID")
	if id == "" {
		response.FailWithMessage("ID必填", c)
		return
	}
	article, err := serviceArticle.GetArticle(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	_ = serviceArticle.IncrementReadCount(id)
	article.ReadCount++
	response.OkWithData(article, c)
}

// GetArticleList 分页获取文章列表
// @Tags     Article
// @Summary  分页获取文章列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query request.ArticleSearch true "分页参数"
// @Success  200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /article/getArticleList [get]
func (a *article) GetArticleList(c *gin.Context) {
	var q request.ArticleSearch
	if err := c.ShouldBindQuery(&q); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceArticle.GetArticleList(q)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     q.Page,
		PageSize: q.PageSize,
	}, "获取成功", c)
}
