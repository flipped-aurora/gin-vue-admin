package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/coffeeModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AddCoffee(c *gin.Context) {
	var coffee coffeeModel.Coffee
	_ = c.ShouldBindJSON(&coffee)
	err := coffee.AddCoffee()
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

func UpdateCoffee(c *gin.Context) {
	var coffee coffeeModel.Coffee
	_ = c.ShouldBindJSON(&coffee)
	err := coffee.UpdateCoffee()
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func DeleteCoffee(c *gin.Context) {
	var cof CoffeeUUID
	var coffee coffeeModel.Coffee
	_ = c.ShouldBindJSON(&cof)
	err := coffee.DeleteCoffee(cof.UUID)
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

func GetCoffeeList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(coffeeModel.Coffee).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"coffeeList": list,
			"total":      total,
			"page":       pageInfo.Page,
			"pageSize":   pageInfo.PageSize,
		})
	}
}

type CoffeeInfo struct {
	PageInfo modelInterface.PageInfo `json:",inline"`
	Code     string                  `json:"code"`
}

func GetCoffeeListByCode(c *gin.Context) {
	var coffeeInfo CoffeeInfo
	_ = c.ShouldBindJSON(&coffeeInfo)
	fmt.Println(coffeeInfo)
	err, list, total := new(coffeeModel.Coffee).GetInfoListByCode(coffeeInfo.PageInfo, coffeeInfo.Code)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"coffeeList": list,
			"total":      total,
			"page":       coffeeInfo.PageInfo.Page,
			"pageSize":   coffeeInfo.PageInfo.PageSize,
		})
	}
}

type SetCoffeeType struct {
	UUID uuid.UUID
	Code string
}

func ChangeCoffeeType(c *gin.Context) {
	var sct SetCoffeeType
	_ = c.ShouldBindJSON(&sct)
	var coffee coffeeModel.Coffee
	err := coffee.UpdateCoffeeType(sct.UUID, sct.Code)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

type CoffeeUUID struct {
	UUID uuid.UUID
}

func GetCoffeeByUUID(c *gin.Context) {
	var cof CoffeeUUID
	var coffee coffeeModel.Coffee
	_ = c.ShouldBindJSON(&cof)
	fmt.Println(cof.UUID)
	err := coffee.GetCoffeeByUUID(cof.UUID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败 %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"coffee": coffee})
	}
}
