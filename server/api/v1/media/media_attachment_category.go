package media

import (
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type AttachmentCategoryApi struct{}

// GetCategoryList
// @Tags      GetCategoryList
// @Summary   媒体库分类列表
// @Security  AttachmentCategory
// @Produce   application/json
// @Success   200   {object}  response.Response{data=media.AttachmentCategory,msg=string}  "媒体库分类列表"
// @Router    /attachmentCategory/getCategoryList [get]
func (a *AttachmentCategoryApi) GetCategoryList(c *gin.Context) {
	res, err := attachmentCategoryService.GetCategoryList(c.Request.Context())
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取分类列表失败!")
		response.FailWithMessage("获取分类列表失败", c)
		return
	}
	response.OkWithData(res, c)
}

// AddCategory
// @Tags      AddCategory
// @Summary   添加媒体库分类
// @Security  AttachmentCategory
// @accept    application/json
// @Produce   application/json
// @Param     data  body      media.AttachmentCategory  true  "媒体库分类数据"// @Success   200   {object}  response.Response{msg=string}   "添加媒体库分类"
// @Router    /attachmentCategory/addCategory [post]
func (a *AttachmentCategoryApi) AddCategory(c *gin.Context) {
	var req media.AttachmentCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("参数错误!")
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := attachmentCategoryService.AddCategory(c.Request.Context(), &req); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("创建/更新失败!")
		response.FailWithMessage("创建/更新失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("创建/更新成功", c)
}

// DeleteCategory
// @Tags      DeleteCategory
// @Summary   删除分类
// @Security  AttachmentCategory
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.GetById                true  "分类id"
// @Success   200   {object}  response.Response{msg=string}  "删除分类"
// @Router    /attachmentCategory/deleteCategory [post]
func (a *AttachmentCategoryApi) DeleteCategory(c *gin.Context) {
	var req common.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if req.ID == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := attachmentCategoryService.DeleteCategory(c.Request.Context(), &req.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
