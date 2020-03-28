package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	// 请自行引入model路径
	"github.com/gin-gonic/gin"
)


// @Tags Test
// @Summary 创建Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "创建Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/createTest [post]
func CreateTest(c *gin.Context) {
	var t autocode.Test
	_ = c.ShouldBindJSON(&t)
	err := t.CreateTest()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}


// @Tags Test
// @Summary 删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /t/deleteTest [post]
func DeleteTest(c *gin.Context) {
	var t autocode.Test
	_ = c.ShouldBindJSON(&t)
	err := t.DeleteTest()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}


// @Tags Test
// @Summary 更新Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "更新Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /t/updateTest [post]
func UpdateTest(c *gin.Context) {
	var t autocode.Test
	_ = c.ShouldBindJSON(&t)
	err,ret := t.UpdateTest()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
		    "ret":ret,
		})
	}
}


// @Tags Test
// @Summary 用id查询Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Test true "用id查询Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /t/findTest [post]
func FindTest(c *gin.Context) {
	var t autocode.Test
	_ = c.ShouldBindJSON(&t)
	err,ret := t.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
		    "ret":ret,
		})
	}
}


// @Tags Test
// @Summary 分页获取Test列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Test列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /t/getTestList [post]
func GetTestList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(autocode.Test).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"autocodeList": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}