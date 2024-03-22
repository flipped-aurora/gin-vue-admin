package webcms

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CateMenusApi struct {
}

var cateMenusService = service.ServiceGroupApp.WebcmsServiceGroup.CateMenusService

// CreateCateMenus 创建CateMenus
// @Tags CateMenus
// @Summary 创建CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.CateMenus true "创建CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cateMenus/createCateMenus [post]
func (cateMenusApi *CateMenusApi) CreateCateMenus(c *gin.Context) {
	var cateMenus webcms.CateMenus
	var cateMenusList []webcms.CateMenus
	err := c.ShouldBindJSON(&cateMenus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMapString("siteinfo")
	// 判断是否批量添加栏目
	nameSplit := strings.Split(cateMenus.Name, "\n")
	for _, v := range nameSplit {
		if v == "" {
			continue
		}
		cateMenus.Name = v
		cateMenus.Siteid = siteinfo["id"]
		cateMenusList = append(cateMenusList, cateMenus)
	}
	// fmt.Println(cateMenusList)
	if err := cateMenusService.CreateCateMenus(cateMenusList); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCateMenus 删除CateMenus
// @Tags CateMenus
// @Summary 删除CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.CateMenus true "删除CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cateMenus/deleteCateMenus [delete]
func (cateMenusApi *CateMenusApi) DeleteCateMenus(c *gin.Context) {
	var cateMenus webcms.CateMenus
	err := c.ShouldBindJSON(&cateMenus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cateMenusService.DeleteCateMenus(cateMenus); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCateMenusByIds 批量删除CateMenus
// @Tags CateMenus
// @Summary 批量删除CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cateMenus/deleteCateMenusByIds [delete]
func (cateMenusApi *CateMenusApi) DeleteCateMenusByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cateMenusService.DeleteCateMenusByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCateMenus 更新CateMenus
// @Tags CateMenus
// @Summary 更新CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.CateMenus true "更新CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cateMenus/updateCateMenus [put]
func (cateMenusApi *CateMenusApi) UpdateCateMenus(c *gin.Context) {
	var cateMenus webcms.CateMenus
	err := c.ShouldBindJSON(&cateMenus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cateMenusService.UpdateCateMenus(cateMenus); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCateMenus 用id查询CateMenus
// @Tags CateMenus
// @Summary 用id查询CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcms.CateMenus true "用id查询CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cateMenus/findCateMenus [get]
func (cateMenusApi *CateMenusApi) FindCateMenus(c *gin.Context) {
	var cateMenus webcms.CateMenus
	err := c.ShouldBindQuery(&cateMenus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recateMenus, err := cateMenusService.GetCateMenus(cateMenus.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recateMenus": recateMenus}, c)
	}
}

// GetCateMenusList 分页获取CateMenus列表
// @Tags CateMenus
// @Summary 分页获取CateMenus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcmsReq.CateMenusSearch true "分页获取CateMenus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cateMenus/getCateMenusList [get]
func (cateMenusApi *CateMenusApi) GetCateMenusList(c *gin.Context) {
	var pageInfo webcmsReq.CateMenusSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMap("siteinfo")
	if list, total, err := cateMenusService.GetCateMenusInfoList(pageInfo, siteinfo); err != nil {
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

// GetTemplateList  获取模板名称列表

func (cateMenusApi *CateMenusApi) GetTemplateList(c *gin.Context) {

	if list, err := cateMenusService.GetTemplateList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

// GetModelsList 获取模型列表
func (cateMenusApi *CateMenusApi) GetModelsList(c *gin.Context) {

	if list, err := cateMenusService.GetModelsList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
