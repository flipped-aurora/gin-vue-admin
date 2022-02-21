package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityBtnApi struct{}

func (a *AuthorityBtnApi) GetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	_ = c.ShouldBindJSON(&req)
	if err, res := authorityBtnService.GetAuthorityBtn(req); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(res, "查询成功", c)
	}
}

func (a *AuthorityBtnApi) SetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	_ = c.ShouldBindJSON(&req)
	if err := authorityBtnService.SetAuthorityBtn(req); err != nil {
		global.GVA_LOG.Error("分配失败!", zap.Error(err))
		response.FailWithMessage("分配失败", c)
	} else {
		response.OkWithMessage("分配成功", c)
	}
}
