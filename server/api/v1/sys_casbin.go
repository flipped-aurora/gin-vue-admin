package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// @Tags casbin
// @Summary 更改角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.CasbinInReceive true "更改角色api权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/UpdateCasbin [post]
func UpdateCasbin(c *gin.Context) {
	var cmr model.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	err := new(model.CasbinModel).UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("添加规则失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "添加规则成功", c)
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
	var cmr model.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	paths := new(model.CasbinModel).GetPolicyPathByAuthorityId(cmr.AuthorityId)
	response.Result(response.SUCCESS, gin.H{"paths": paths}, "获取规则成功", c)
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
	response.Result(response.SUCCESS, gin.H{"pathParam": pathParam, "query": query}, "获取规则成功", c)
}
