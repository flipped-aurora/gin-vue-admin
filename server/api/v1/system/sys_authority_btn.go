package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityBtnApi struct{}

// GetAuthorityBtn
// @Tags      AuthorityBtn
// @Summary   获取权限按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SysAuthorityBtnReq                                      true  "菜单id, 角色id, 选中的按钮id"
// @Success   200   {object}  response.Response{data=response.SysAuthorityBtnRes,msg=string}  "返回列表成功"
// @Router    /authorityBtn/getAuthorityBtn [post]
func (a *AuthorityBtnApi) GetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := authorityBtnService.GetAuthorityBtn(req)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFailErr"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithDetailed(res, global.Translate("general.querySuccess"), c)
}

// SetAuthorityBtn
// @Tags      AuthorityBtn
// @Summary   设置权限按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SysAuthorityBtnReq     true  "菜单id, 角色id, 选中的按钮id"
// @Success   200   {object}  response.Response{msg=string}  "返回列表成功"
// @Router    /authorityBtn/setAuthorityBtn [post]
func (a *AuthorityBtnApi) SetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityBtnService.SetAuthorityBtn(req)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_authority.assignmentFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_authority.assignmentFail"), c)
		return
	}
	response.OkWithMessage(global.Translate("sys_authority.assignmentSuccess"), c)
}

// CanRemoveAuthorityBtn
// @Tags      AuthorityBtn
// @Summary   设置权限按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "删除成功"
// @Router    /authorityBtn/canRemoveAuthorityBtn [post]
func (a *AuthorityBtnApi) CanRemoveAuthorityBtn(c *gin.Context) {
	id := c.Query("id")
	err := authorityBtnService.CanRemoveAuthorityBtn(id)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}
