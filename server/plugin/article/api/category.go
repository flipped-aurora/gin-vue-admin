package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Category = new(category)

type category struct{}

// CreateCategory 创建分类
// @Tags     ArticleCategory
// @Summary  创建文章分类
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body model.Category true "分类信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /article/createCategory [post]
func (a *category) CreateCategory(c *gin.Context) {
	var cat model.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serviceCategory.CreateCategory(&cat); err != nil {
		global.GVA_LOG.Error("创建分类失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCategory 删除分类
// @Tags     ArticleCategory
// @Summary  删除文章分类
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    ID query string true "分类ID"
// @Success  200 {object} response.Response{msg=string} "删除成功"
// @Router   /article/deleteCategory [delete]
func (a *category) DeleteCategory(c *gin.Context) {
	id := c.Query("ID")
	if id == "" {
		response.FailWithMessage("ID必填", c)
		return
	}
	if err := serviceCategory.DeleteCategory(id); err != nil {
		global.GVA_LOG.Error("删除分类失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateCategory 更新分类
// @Tags     ArticleCategory
// @Summary  更新文章分类
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body model.Category true "分类信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /article/updateCategory [put]
func (a *category) UpdateCategory(c *gin.Context) {
	var cat model.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serviceCategory.UpdateCategory(cat); err != nil {
		global.GVA_LOG.Error("更新分类失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetCategoryList 分页获取分类列表
// @Tags     ArticleCategory
// @Summary  分页获取分类列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query request.CategorySearch true "分页参数"
// @Success  200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /article/getCategoryList [get]
func (a *category) GetCategoryList(c *gin.Context) {
	var q request.CategorySearch
	if err := c.ShouldBindQuery(&q); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceCategory.GetCategoryList(q)
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

// GetAllCategories 获取所有分类(不分页,下拉用)
// @Tags     ArticleCategory
// @Summary  获取所有分类
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200 {object} response.Response{data=[]model.Category,msg=string} "获取成功"
// @Router   /article/getAllCategories [get]
func (a *category) GetAllCategories(c *gin.Context) {
	list, err := serviceCategory.GetAllCategories()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(list, c)
}
