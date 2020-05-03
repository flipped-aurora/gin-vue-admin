package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
)

// @Tags SysApi
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func CreateTemp(c *gin.Context) {
	var a model.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if a.AutoCreateApiToSql {
		apiList := [5]model.SysApi{
			{
				Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
				Description: "新增" + a.StructName,
				ApiGroup:    a.Abbreviation,
				Method:      "POST",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
				Description: "删除" + a.StructName,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
				Description: "更新" + a.StructName,
				ApiGroup:    a.Abbreviation,
				Method:      "POST",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
				Description: "根据ID获取" + a.StructName,
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
				Description: "获取" + a.StructName + "列表",
				ApiGroup:    a.StructName,
				Method:      "GET",
			},
		}
		for _, v := range apiList {
			errC := service.CreateApi(v)
			if errC != nil {
				c.Writer.Header().Add("success", "false")
				c.Writer.Header().Add("msg", url.QueryEscape(fmt.Sprintf("自动化创建失败，%v，请自行清空垃圾数据", errC)))
				return
			}
		}
	}
	err := service.CreateTemp(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		os.Remove("./ginvueadmin.zip")
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		os.Remove("./ginvueadmin.zip")
	}
}
