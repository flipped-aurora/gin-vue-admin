package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/controller/servers"
	"main/middleware"
	"main/model/dbModel"
)

// @Tags Menu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "可以什么都不填"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.CustomClaims)
	err, menus := new(dbModel.Menu).GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{"menus": menus})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"menus": menus})
	}
}
