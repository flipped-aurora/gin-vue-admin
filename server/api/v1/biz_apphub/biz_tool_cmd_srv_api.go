package biz_apphub

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	biz_apphub1 "github.com/flipped-aurora/gin-vue-admin/server/service/biz_apphub"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizToolCmdSrvApiApi struct{}

// CreateBizToolCmdSrvApi 创建后端工具指令api
// @Tags BizToolCmdSrvApi
// @Summary 创建后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizToolCmdSrvApi true "创建后端工具指令api"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bizToolCmdSrvApi/createBizToolCmdSrvApi [post]
func (bizToolCmdSrvApiApi *BizToolCmdSrvApiApi) CreateBizToolCmdSrvApi(c *gin.Context) {
	var (
		err              error
		bizToolCmdSrvApi biz_apphub.BizToolCmdSrvApi
		installInfo      *biz_apphub1.InstallInfo
	)
	defer global.GVA_LOG.Info(fmt.Sprintf("CreateBizToolCmdSrvApi err:%v req:%+v installInfo:%+v ", err, bizToolCmdSrvApi, installInfo))
	err = c.ShouldBindJSON(&bizToolCmdSrvApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizToolCmdSrvApi.CreatedBy = utils.GetUserID(c)
	bizToolCmdSrvApi.OperateUser = c.GetString("user")
	installInfo, err = bizToolCmdSrvApiService.Install(bizToolCmdSrvApi) //安装软件
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//host := "https://apphub.geeleo.com/api"
	//if os.Getenv("env") == "local" {
	//	host = "http://127.0.0.1:8888"
	//}
	bizToolCmdSrvApi.ApiPath = fmt.Sprintf("/bizAppHub/api/cmd/call/%s/%s",
		bizToolCmdSrvApi.OperateUser, bizToolCmdSrvApi.AppCode)

	err = bizToolCmdSrvApiService.CreateBizToolCmdSrvApi(&bizToolCmdSrvApi)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizToolCmdSrvApi 删除后端工具指令api
// @Tags BizToolCmdSrvApi
// @Summary 删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizToolCmdSrvApi true "删除后端工具指令api"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bizToolCmdSrvApi/deleteBizToolCmdSrvApi [delete]
func (bizToolCmdSrvApiApi *BizToolCmdSrvApiApi) DeleteBizToolCmdSrvApi(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := bizToolCmdSrvApiService.DeleteBizToolCmdSrvApi(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizToolCmdSrvApiByIds 批量删除后端工具指令api
// @Tags BizToolCmdSrvApi
// @Summary 批量删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bizToolCmdSrvApi/deleteBizToolCmdSrvApiByIds [delete]
func (bizToolCmdSrvApiApi *BizToolCmdSrvApiApi) DeleteBizToolCmdSrvApiByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := bizToolCmdSrvApiService.DeleteBizToolCmdSrvApiByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizToolCmdSrvApi 更新后端工具指令api
// @Tags BizToolCmdSrvApi
// @Summary 更新后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizToolCmdSrvApi true "更新后端工具指令api"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizToolCmdSrvApi/updateBizToolCmdSrvApi [put]
func (bizToolCmdSrvApiApi *BizToolCmdSrvApiApi) UpdateBizToolCmdSrvApi(c *gin.Context) {
	var bizToolCmdSrvApi biz_apphub.BizToolCmdSrvApi
	err := c.ShouldBindJSON(&bizToolCmdSrvApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizToolCmdSrvApi.UpdatedBy = utils.GetUserID(c)
	err = bizToolCmdSrvApiService.UpdateBizToolCmdSrvApi(bizToolCmdSrvApi)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizToolCmdSrvApi 用id查询后端工具指令api
// @Tags BizToolCmdSrvApi
// @Summary 用id查询后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphub.BizToolCmdSrvApi true "用id查询后端工具指令api"
// @Success 200 {object} response.Response{data=biz_apphub.BizToolCmdSrvApi,msg=string} "查询成功"
// @Router /bizToolCmdSrvApi/findBizToolCmdSrvApi [get]
func (bizToolCmdSrvApiApi *BizToolCmdSrvApiApi) FindBizToolCmdSrvApi(c *gin.Context) {
	ID := c.Query("ID")
	rebizToolCmdSrvApi, err := bizToolCmdSrvApiService.GetBizToolCmdSrvApi(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebizToolCmdSrvApi, c)
}

// GetBizToolCmdSrvApiList 分页获取后端工具指令api列表
// @Tags BizToolCmdSrvApi
// @Summary 分页获取后端工具指令api列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizToolCmdSrvApiSearch true "分页获取后端工具指令api列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bizToolCmdSrvApi/getBizToolCmdSrvApiList [get]
func (bizToolCmdSrvApiApi *BizToolCmdSrvApiApi) GetBizToolCmdSrvApiList(c *gin.Context) {
	var pageInfo biz_apphubReq.BizToolCmdSrvApiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bizToolCmdSrvApiService.GetBizToolCmdSrvApiInfoList(pageInfo)
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

// GetBizToolCmdSrvApiPublic 不需要鉴权的后端工具指令api接口
// @Tags BizToolCmdSrvApi
// @Summary 不需要鉴权的后端工具指令api接口
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizToolCmdSrvApiSearch true "分页获取后端工具指令api列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bizToolCmdSrvApi/getBizToolCmdSrvApiPublic [get]
func (bizToolCmdSrvApiApi *BizToolCmdSrvApiApi) GetBizToolCmdSrvApiPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的后端工具指令api接口信息",
	}, "获取成功", c)
}
