package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/oss"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizAppHubApi struct{}

// CreateBizAppHub 创建biz_apphub
// @Tags BizAppHub
// @Summary 创建biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "创建biz_apphub"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bizAppHub/createBizAppHub [post]
func (bizAppHubApi *BizAppHubApi) CreateBizAppHub(c *gin.Context) {
	var bizAppHub biz_apphub.BizAppHub
	err := c.ShouldBindJSON(&bizAppHub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizAppHub.CreatedBy = utils.GetUserID(c)
	err = bizAppHubService.CreateBizAppHub(&bizAppHub)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizAppHub 删除biz_apphub
// @Tags BizAppHub
// @Summary 删除biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "删除biz_apphub"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bizAppHub/deleteBizAppHub [delete]
func (bizAppHubApi *BizAppHubApi) DeleteBizAppHub(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := bizAppHubService.DeleteBizAppHub(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizAppHubByIds 批量删除biz_apphub
// @Tags BizAppHub
// @Summary 批量删除biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bizAppHub/deleteBizAppHubByIds [delete]
func (bizAppHubApi *BizAppHubApi) DeleteBizAppHubByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := bizAppHubService.DeleteBizAppHubByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizAppHub 更新biz_apphub
// @Tags BizAppHub
// @Summary 更新biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "更新biz_apphub"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizAppHub/updateBizAppHub [put]
func (bizAppHubApi *BizAppHubApi) UpdateBizAppHub(c *gin.Context) {
	var bizAppHub biz_apphub.BizAppHub
	err := c.ShouldBindJSON(&bizAppHub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizAppHub.UpdatedBy = utils.GetUserID(c)
	err = bizAppHubService.UpdateBizAppHub(bizAppHub)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizAppHub 用id查询biz_apphub
// @Tags BizAppHub
// @Summary 用id查询biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphub.BizAppHub true "用id查询biz_apphub"
// @Success 200 {object} response.Response{data=biz_apphub.BizAppHub,msg=string} "查询成功"
// @Router /bizAppHub/findBizAppHub [get]
func (bizAppHubApi *BizAppHubApi) FindBizAppHub(c *gin.Context) {
	ID := c.Query("ID")
	rebizAppHub, err := bizAppHubService.GetBizAppHub(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebizAppHub, c)
}

// GetBizAppHubList 分页获取biz_apphub列表
// @Tags BizAppHub
// @Summary 分页获取biz_apphub列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizAppHubSearch true "分页获取biz_apphub列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bizAppHub/getBizAppHubList [get]
func (bizAppHubApi *BizAppHubApi) GetBizAppHubList(c *gin.Context) {
	var pageInfo biz_apphubReq.BizAppHubSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bizAppHubService.GetBizAppHubInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetBizAppHubPublic 不需要鉴权的biz_apphub接口
// @Tags BizAppHub
// @Summary 不需要鉴权的biz_apphub接口
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizAppHubSearch true "分页获取biz_apphub列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bizAppHub/getBizAppHubPublic [get]
func (bizAppHubApi *BizAppHubApi) GetBizAppHubPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的biz_apphub接口信息",
	}, "获取成功", c)
}

func (bizAppHubApi *BizAppHubApi) GetUploadToken(c *gin.Context) {
	//token := service.NewDefaultService().Oss.GetUploadToken()
	store := oss.NewDefaultQiNiu()
	response.OkWithData(gin.H{"token": store.GetUploadToken()}, c)
}
