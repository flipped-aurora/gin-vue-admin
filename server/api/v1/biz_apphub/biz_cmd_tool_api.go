package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizCmdToolApiApi struct{}

// CreateBizCmdToolApi 创建后端工具指令api
// @Tags BizCmdToolApi
// @Summary 创建后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCmdToolApi true "创建后端工具指令api"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bizCmdToolApi/createBizCmdToolApi [post]
func (bizCmdToolApiApi *BizCmdToolApiApi) CreateBizCmdToolApi(c *gin.Context) {
	var bizCmdToolApi biz_apphub.BizCmdToolApi
	err := c.ShouldBindJSON(&bizCmdToolApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizCmdToolApi.CreatedBy = utils.GetUserID(c)
	err = bizCmdToolApiService.CreateBizCmdToolApi(&bizCmdToolApi)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizCmdToolApi 删除后端工具指令api
// @Tags BizCmdToolApi
// @Summary 删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCmdToolApi true "删除后端工具指令api"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bizCmdToolApi/deleteBizCmdToolApi [delete]
func (bizCmdToolApiApi *BizCmdToolApiApi) DeleteBizCmdToolApi(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := bizCmdToolApiService.DeleteBizCmdToolApi(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizCmdToolApiByIds 批量删除后端工具指令api
// @Tags BizCmdToolApi
// @Summary 批量删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bizCmdToolApi/deleteBizCmdToolApiByIds [delete]
func (bizCmdToolApiApi *BizCmdToolApiApi) DeleteBizCmdToolApiByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := bizCmdToolApiService.DeleteBizCmdToolApiByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizCmdToolApi 更新后端工具指令api
// @Tags BizCmdToolApi
// @Summary 更新后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCmdToolApi true "更新后端工具指令api"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizCmdToolApi/updateBizCmdToolApi [put]
func (bizCmdToolApiApi *BizCmdToolApiApi) UpdateBizCmdToolApi(c *gin.Context) {
	var bizCmdToolApi biz_apphub.BizCmdToolApi
	err := c.ShouldBindJSON(&bizCmdToolApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizCmdToolApi.UpdatedBy = utils.GetUserID(c)
	err = bizCmdToolApiService.UpdateBizCmdToolApi(bizCmdToolApi)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizCmdToolApi 用id查询后端工具指令api
// @Tags BizCmdToolApi
// @Summary 用id查询后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphub.BizCmdToolApi true "用id查询后端工具指令api"
// @Success 200 {object} response.Response{data=biz_apphub.BizCmdToolApi,msg=string} "查询成功"
// @Router /bizCmdToolApi/findBizCmdToolApi [get]
func (bizCmdToolApiApi *BizCmdToolApiApi) FindBizCmdToolApi(c *gin.Context) {
	ID := c.Query("ID")
	rebizCmdToolApi, err := bizCmdToolApiService.GetBizCmdToolApi(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebizCmdToolApi, c)
}

// GetBizCmdToolApiList 分页获取后端工具指令api列表
// @Tags BizCmdToolApi
// @Summary 分页获取后端工具指令api列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizCmdToolApiSearch true "分页获取后端工具指令api列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bizCmdToolApi/getBizCmdToolApiList [get]
func (bizCmdToolApiApi *BizCmdToolApiApi) GetBizCmdToolApiList(c *gin.Context) {
	var pageInfo biz_apphubReq.BizCmdToolApiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bizCmdToolApiService.GetBizCmdToolApiInfoList(pageInfo)
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

// GetBizCmdToolApiPublic 不需要鉴权的后端工具指令api接口
// @Tags BizCmdToolApi
// @Summary 不需要鉴权的后端工具指令api接口
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizCmdToolApiSearch true "分页获取后端工具指令api列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bizCmdToolApi/getBizCmdToolApiPublic [get]
func (bizCmdToolApiApi *BizCmdToolApiApi) GetBizCmdToolApiPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的后端工具指令api接口信息",
	}, "获取成功", c)
}
