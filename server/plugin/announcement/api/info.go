package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Info = new(info)

type info struct{}

// CreateInfo 创建公告
// @Tags Info
// @Summary 创建公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "创建公告"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /info/createInfo [post]
func (a *info) CreateInfo(c *gin.Context) {
	var info model.Info
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceInfo.CreateInfo(&info)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.createSuccss"), c)
}

// DeleteInfo 删除公告
// @Tags Info
// @Summary 删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "删除公告"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /info/deleteInfo [delete]
func (a *info) DeleteInfo(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceInfo.DeleteInfo(ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFail"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteInfoByIds 批量删除公告
// @Tags Info
// @Summary 批量删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /info/deleteInfoByIds [delete]
func (a *info) DeleteInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := serviceInfo.DeleteInfoByIds(IDs); err != nil {
		global.GVA_LOG.Error(global.Translate("sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_operation_record.batchDeleteFail"), c)
		return
	}
	response.OkWithMessage(global.Translate("sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateInfo 更新公告
// @Tags Info
// @Summary 更新公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "更新公告"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /info/updateInfo [put]
func (a *info) UpdateInfo(c *gin.Context) {
	var info model.Info
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceInfo.UpdateInfo(info)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFail"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindInfo 用id查询公告
// @Tags Info
// @Summary 用id查询公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Info true "用id查询公告"
// @Success 200 {object} response.Response{data=model.Info,msg=string} "查询成功"
// @Router /info/findInfo [get]
func (a *info) FindInfo(c *gin.Context) {
	ID := c.Query("ID")
	reinfo, err := serviceInfo.GetInfo(ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFail"), c)
		return
	}
	response.OkWithData(reinfo, c)
}

// GetInfoList 分页获取公告列表
// @Tags Info
// @Summary 分页获取公告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.InfoSearch true "分页获取公告列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /info/getInfoList [get]
func (a *info) GetInfoList(c *gin.Context) {
	var pageInfo request.InfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceInfo.GetInfoInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetInfoDataSource 获取Info的数据源
// @Tags Info
// @Summary 获取Info的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /info/getInfoDataSource [get]
func (a *info) GetInfoDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	dataSource, err := serviceInfo.GetInfoDataSource()
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFail"), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetInfoPublic 不需要鉴权的公告接口
// @Tags Info
// @Summary 不需要鉴权的公告接口
// @accept application/json
// @Produce application/json
// @Param data query request.InfoSearch true "分页获取公告列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /info/getInfoPublic [get]
func (a *info) GetInfoPublic(c *gin.Context) {
	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{"info": global.Translate("announcement.publicAnnouncementAPI")}, global.Translate("general.getDataSuccess"), c)
}
