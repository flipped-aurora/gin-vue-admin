package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type certCategory struct{}

// CreateCertCategory 创建项目
func (a *certCategory) CreateCertCategory(c *gin.Context) {
	var cc model.CertCategory
	_ = c.ShouldBindJSON(&cc)
	if err := serviceCertCategory.CreateCertCategory(cc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCertCategory 删除项目
func (a *certCategory) DeleteCertCategory(c *gin.Context) {
	var cc model.CertCategory
	_ = c.ShouldBindJSON(&cc)
	if err := serviceCertCategory.DeleteCertCategory(cc); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateCertCategory 更新项目
func (a *certCategory) UpdateCertCategory(c *gin.Context) {
	var cc model.CertCategory
	_ = c.ShouldBindJSON(&cc)
	if err := serviceCertCategory.UpdateCertCategory(cc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetCertCategoryList 获取项目列表
func (a *certCategory) GetCertCategoryList(c *gin.Context) {
	var pageInfo request.CertCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := serviceCertCategory.GetCertCategoryList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// BatchUpdateCertCategory 批量修改项目
func (a *certCategory) BatchUpdateCertCategory(c *gin.Context) {
	var info request.BatchUpdateCategory
	_ = c.ShouldBindJSON(&info)
	if err := serviceCertCategory.BatchUpdateCertCategory(info); err != nil {
		global.GVA_LOG.Error("批量修改失败!", zap.Error(err))
		response.FailWithMessage("批量修改失败", c)
	} else {
		response.OkWithMessage("批量修改成功", c)
	}
}
