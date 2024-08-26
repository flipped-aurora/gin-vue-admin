package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service/biz_apphub"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetDeployList 分页获取biz_apphub_record列表
// @Tags BizAppHub
// @Summary 分页获取biz_apphub_record列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizAppHubSearch true "分页获取biz_apphub列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bizAppHub/getDeployList [get]
func (bizAppHubApi *BizAppHubApi) GetDeployList(c *gin.Context) {
	var pageInfo biz_apphubReq.GetDeployList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.OperateUser = c.GetString("user")
	list, total, err := bizAppHubService.GetDeployList(pageInfo)
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

// RollbackVersion 回滚版本biz_apphub
// @Tags BizAppHub
// @Summary 回滚版本biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "回滚版本biz_apphub"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizAppHub/rollbackVersion [put]
func (bizAppHubApi *BizAppHubApi) RollbackVersion(c *gin.Context) {
	var bizAppHub biz_apphubReq.RollbackVersion
	err := c.ShouldBindJSON(&bizAppHub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//bizAppHub.OperateUser=c.GetString("user")
	//bizAppHub.UpdatedBy = utils.GetUserID(c)
	bizAppHub.OperateUser = c.GetString("user")
	err = bizAppHubService.RollbackVersion(bizAppHub)
	if err != nil {
		global.GVA_LOG.Error("回滚版本失败!", zap.Error(err))
		response.FailWithMessage("回滚版本失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("回滚版本成功", c)
}

// Call 后端命令接口
// @Tags BizAppHub
// @Summary 后端命令接口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "后端命令接口"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizAppHub/api/caller [post]
func (bizAppHubApi *BizAppHubApi) Call(c *gin.Context) {
	//var bizAppHub biz_apphubReq.RollbackVersion
	var req biz_apphubReq.Call
	//var req request.
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//bizAppHub.OperateUser=c.GetString("user")
	//bizAppHub.UpdatedBy = utils.GetUserID(c)
	//bizAppHub.OperateUser = c.GetString("user")
	//todo 验证用户是否有调用该应用的权限
	j, err := req.RequestJSON()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.Req = j
	call, err := biz_apphub.NewCaller("").Call(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	c.Data(200, "application/json", []byte(call))
}
