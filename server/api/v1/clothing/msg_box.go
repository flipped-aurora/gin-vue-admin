package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MsgBoxApi struct {
}

var msgBoxService = service.ServiceGroupApp.ClothingServiceGroup.MsgBoxService

func (msgBoxApi *MsgBoxApi) CreateMsgBox(c *gin.Context) {
	var msgBox clothing.MsgBox
	err := c.ShouldBindJSON(&msgBox)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msgBox.CreatedBy = utils.GetUserID(c)
	if err := msgBoxService.CreateMsgBox(&msgBox); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (msgBoxApi *MsgBoxApi) DeleteMsgBox(c *gin.Context) {
	var msgBox clothing.MsgBox
	err := c.ShouldBindJSON(&msgBox)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msgBox.DeletedBy = utils.GetUserID(c)
	if err := msgBoxService.DeleteMsgBox(msgBox); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (msgBoxApi *MsgBoxApi) DeleteMsgBoxByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := msgBoxService.DeleteMsgBoxByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (msgBoxApi *MsgBoxApi) UpdateMsgBox(c *gin.Context) {
	var msgBox clothing.MsgBox
	err := c.ShouldBindJSON(&msgBox)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msgBox.UpdatedBy = utils.GetUserID(c)
	if err := msgBoxService.UpdateMsgBox(msgBox); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (msgBoxApi *MsgBoxApi) FindMsgBox(c *gin.Context) {
	var msgBox clothing.MsgBox
	err := c.ShouldBindQuery(&msgBox)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remsgBox, err := msgBoxService.GetMsgBox(msgBox.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remsgBox": remsgBox}, c)
	}
}

func (msgBoxApi *MsgBoxApi) GetMsgBoxList(c *gin.Context) {
	var pageInfo clothingReq.MsgBoxSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := msgBoxService.GetMsgBoxInfoList(pageInfo); err != nil {
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

// GetMsgBoxList 分页获取收到的信息
// @Tags MsgBox
// @Summary 分页获取收到的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.MsgBoxSearch true "分页获取收到的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgBox/getMyMsgList [get]
func (msgBoxApi *MsgBoxApi) GetMyMsgList(c *gin.Context) {
	var pageInfo clothingReq.MsgBoxSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.To = utils.GetUserID(c)
	if list, total, err := msgBoxService.GetMsgBoxInfoList(pageInfo); err != nil {
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

// GetMySendMsgList 分页获取发送的信息
// @Tags MsgBox
// @Summary 分页获取发送的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.MsgBoxSearch true "分页获取发送的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgBox/getMySendMsgList [get]
func (msgBoxApi *MsgBoxApi) GetMySendMsgList(c *gin.Context) {
	var pageInfo clothingReq.MsgBoxSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.From = utils.GetUserID(c)
	if list, total, err := msgBoxService.GetMsgBoxInfoList(pageInfo); err != nil {
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

// SetRead 设置已读
// @Tags MsgBox
// @Summary 设置已读
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.MsgBox true "设置已读"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msgBox/setRead [get]
func (msgBoxApi *MsgBoxApi) SetRead(c *gin.Context) {
	var msgBox clothing.MsgBox
	err := c.ShouldBindQuery(&msgBox)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if msg, err := msgBoxService.GetMsgBox(msgBox.ID); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		if msg.To == utils.GetUserID(c) {
			if err := msgBoxService.SetRead(msgBox.ID); err == nil {
				response.Ok(c)
				return
			}
		}
		response.FailWithMessage("操作失败", c)
		return
	}
}
