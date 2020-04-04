package api

import (
	"fmt"
	// 请自行引入model路径
	"github.com/gin-gonic/gin"
)


// @Tags {{.StructName}}
// @Summary 创建{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "创建{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func Create{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.PackageName}}.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	err := {{.Abbreviation}}.Create{{.StructName}}()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}


// @Tags {{.StructName}}
// @Summary 删除{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "删除{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [post]
func Delete{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.PackageName}}.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	err := {{.Abbreviation}}.Delete{{.StructName}}()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}


// @Tags {{.StructName}}
// @Summary 更新{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "更新{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [post]
func Update{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.PackageName}}.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	err,re{{.Abbreviation}} := {{.Abbreviation}}.Update{{.StructName}}()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
		    "re{{.Abbreviation}}":re{{.Abbreviation}},
		})
	}
}


// @Tags {{.StructName}}
// @Summary 用id查询{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "用id查询{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [post]
func Find{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.PackageName}}.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	err,re{{.Abbreviation}} := {{.Abbreviation}}.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
		    "re{{.Abbreviation}}":re{{.Abbreviation}},
		})
	}
}


// @Tags {{.StructName}}
// @Summary 分页获取{{.StructName}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取{{.StructName}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [post]
func Get{{.StructName}}List(c *gin.Context) {
	var pageInfo model.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new({{.PackageName}}.{{.StructName}}).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"{{.PackageName}}List": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}