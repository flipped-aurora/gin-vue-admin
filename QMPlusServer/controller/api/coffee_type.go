package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/coffeeModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
)

func AddCoffeeType(c *gin.Context) {
	var coffeeType coffeeModel.CoffeeType
	_ = c.ShouldBindJSON(&coffeeType)
	err := coffeeType.AddCoffeeType()
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

func UpdateCoffeeType(c *gin.Context) {
	var coffeeType coffeeModel.CoffeeType
	_ = c.ShouldBindJSON(&coffeeType)
	err := coffeeType.UpdateCoffeeType()
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func DeleteCoffeeType(c *gin.Context) {
	var coffeeType coffeeModel.CoffeeType
	var coffeeTypeCode CoffeeTypeCode
	err := coffeeType.DeleteCoffeeType(coffeeTypeCode.Code)
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

func GetCoffeeTypeList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(coffeeModel.CoffeeType).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"coffeetype": list,
			"total":      total,
			"page":       pageInfo.Page,
			"pageSize":   pageInfo.PageSize,
		})
	}
}

type CoffeeTypeCode struct {
	Code string
}

func GetCoffeeByCode(c *gin.Context) {
	var code CoffeeTypeCode
	var cty coffeeModel.CoffeeType
	_ = c.ShouldBindJSON(&code)
	err := cty.GetCoffeeTypeByCode(code.Code)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败,%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"coffeetype": cty})
	}
}
