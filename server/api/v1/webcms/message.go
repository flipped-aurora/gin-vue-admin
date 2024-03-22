package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MessageApi struct {
}

var messageService = service.ServiceGroupApp.WebcmsServiceGroup.MessageService

// CreateMessage 创建Message
// @Tags Message
// @Summary 创建Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Message true "创建Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /message/createMessage [post]
func (messageApi *MessageApi) CreateMessage(c *gin.Context) {
	var message webcms.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":    {utils.NotEmpty()},
		"Email":   {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
	}
	if err := utils.Verify(message, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := messageService.CreateMessage(message); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMessage 删除Message
// @Tags Message
// @Summary 删除Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Message true "删除Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /message/deleteMessage [delete]
func (messageApi *MessageApi) DeleteMessage(c *gin.Context) {
	var message webcms.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := messageService.DeleteMessage(message); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMessageByIds 批量删除Message
// @Tags Message
// @Summary 批量删除Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /message/deleteMessageByIds [delete]
func (messageApi *MessageApi) DeleteMessageByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := messageService.DeleteMessageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMessage 更新Message
// @Tags Message
// @Summary 更新Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Message true "更新Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /message/updateMessage [put]
func (messageApi *MessageApi) UpdateMessage(c *gin.Context) {
	var message webcms.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":    {utils.NotEmpty()},
		"Email":   {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
	}
	if err := utils.Verify(message, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := messageService.UpdateMessage(message); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMessage 用id查询Message
// @Tags Message
// @Summary 用id查询Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcms.Message true "用id查询Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /message/findMessage [get]
func (messageApi *MessageApi) FindMessage(c *gin.Context) {
	var message webcms.Message
	err := c.ShouldBindQuery(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remessage, err := messageService.GetMessage(message.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remessage": remessage}, c)
	}
}

// GetMessageList 分页获取Message列表
// @Tags Message
// @Summary 分页获取Message列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcmsReq.MessageSearch true "分页获取Message列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /message/getMessageList [get]
func (messageApi *MessageApi) GetMessageList(c *gin.Context) {
	var pageInfo webcmsReq.MessageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := messageService.GetMessageInfoList(pageInfo); err != nil {
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
