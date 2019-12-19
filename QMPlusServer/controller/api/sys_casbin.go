package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

type CasbinInReceive struct {
	AuthorityId string   `json:"authorityId"`
	Paths       []string `json:paths`
}

// @Tags casbin
// @Summary 更改角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "更改角色api权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/casbinPUpdata [post]
func CasbinPUpdata(c *gin.Context) {
	var cmr CasbinInReceive
	_ = c.ShouldBind(&cmr)
	err := new(sysModel.CasbinModel).CasbinPUpdata(cmr.AuthorityId, cmr.Paths)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加规则失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加规则成功", gin.H{})
	}
}

// @Tags casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func GetPolicyPathByAuthorityId(c *gin.Context) {
	var cmr CasbinInReceive
	_ = c.ShouldBind(&cmr)
	paths := new(sysModel.CasbinModel).GetPolicyPathByAuthorityId(cmr.AuthorityId)
	servers.ReportFormat(c, true, "获取规则成功", gin.H{"paths": paths})
}
