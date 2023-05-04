package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MsgBoxApi struct {
}

var msgBoxService = service.ServiceGroupApp.ClothingServiceGroup.MsgBoxService


// CreateMsgBox 创建MsgBox
// @Tags MsgBox
// @Summary 创建MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.MsgBox true "创建MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgBox/createMsgBox [post]
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

// DeleteMsgBox 删除MsgBox
// @Tags MsgBox
// @Summary 删除MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.MsgBox true "删除MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msgBox/deleteMsgBox [delete]
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

// DeleteMsgBoxByIds 批量删除MsgBox
// @Tags MsgBox
// @Summary 批量删除MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /msgBox/deleteMsgBoxByIds [delete]
func (msgBoxApi *MsgBoxApi) DeleteMsgBoxByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := msgBoxService.DeleteMsgBoxByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMsgBox 更新MsgBox
// @Tags MsgBox
// @Summary 更新MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.MsgBox true "更新MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msgBox/updateMsgBox [put]
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

// FindMsgBox 用id查询MsgBox
// @Tags MsgBox
// @Summary 用id查询MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.MsgBox true "用id查询MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msgBox/findMsgBox [get]
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

// GetMsgBoxList 分页获取MsgBox列表
// @Tags MsgBox
// @Summary 分页获取MsgBox列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.MsgBoxSearch true "分页获取MsgBox列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgBox/getMsgBoxList [get]
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
