package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

// @Tags casbin
// @Summary 更改角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.CasbinInReceive true "更改角色api权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/casbinPUpdata [post]
func CasbinPUpdata(c *gin.Context) {
	var cmr sysModel.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	err := new(sysModel.CasbinModel).CasbinPUpdata(cmr.AuthorityId, cmr.CasbinInfos)
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
// @Param data body sysModel.CasbinInReceive true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func GetPolicyPathByAuthorityId(c *gin.Context) {
	var cmr sysModel.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	paths := new(sysModel.CasbinModel).GetPolicyPathByAuthorityId(cmr.AuthorityId)
	servers.ReportFormat(c, true, "获取规则成功", gin.H{"paths": paths})
}

// @Tags casbin
// @Summary casb RBAC RESTFUL测试路由
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.CasbinInReceive true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/CasbinTest [get]
func CasbinTest(c *gin.Context) {
	// 测试restful以及占位符代码  随意书写
	pathParam := c.Param("pathParam")
	query := c.Query("query")
	servers.ReportFormat(c, true, "获取规则成功", gin.H{"pathParam": pathParam, "query": query})
}
