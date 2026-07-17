package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	sysReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type ApiTokenApi struct{}

// CreateApiToken 签发Token
// @Tags      SysApiToken
// @Summary   签发ApiToken
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object                                       true  "用户ID, 角色ID, 有效天数(-1为永久), 备注"
// @Success   200   {object}  response.Response{data=object,msg=string}    "签发ApiToken"
// @Router    /sysApiToken/createApiToken [post]
func (s *ApiTokenApi) CreateApiToken(c *gin.Context) {
	var req struct {
		UserID      uint   `json:"userId"`
		AuthorityID uint   `json:"authorityId"`
		Days        int    `json:"days"` // -1为永久, 其他为天数
		Remark      string `json:"remark"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	token := system.SysApiToken{
		UserID:      req.UserID,
		AuthorityID: req.AuthorityID,
		Remark:      req.Remark,
	}

	jwtStr, err := apiTokenService.CreateApiToken(c.Request.Context(), token, req.Days)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("签发失败!")
		response.FailWithMessage("签发失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"token": jwtStr}, "签发成功", c)
}

// GetApiTokenList 获取列表
// @Tags      SysApiToken
// @Summary   分页获取ApiToken列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      sysReq.SysApiTokenSearch                            true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取ApiToken列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysApiToken/getApiTokenList [post]
func (s *ApiTokenApi) GetApiTokenList(c *gin.Context) {
	var pageInfo sysReq.SysApiTokenSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := apiTokenService.GetApiTokenList(c.Request.Context(), pageInfo)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// DeleteApiToken 作废Token
// @Tags      SysApiToken
// @Summary   作废ApiToken
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApiToken             true  "ApiToken的ID"
// @Success   200   {object}  response.Response{msg=string}  "作废ApiToken"
// @Router    /sysApiToken/deleteApiToken [post]
func (s *ApiTokenApi) DeleteApiToken(c *gin.Context) {
	var req system.SysApiToken
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiTokenService.DeleteApiToken(c.Request.Context(), req.ID)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("作废失败!")
		response.FailWithMessage("作废失败", c)
		return
	}
	response.OkWithMessage("作废成功", c)
}
