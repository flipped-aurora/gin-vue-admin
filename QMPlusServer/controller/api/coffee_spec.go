package api

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/coffeeModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CoffeeId struct {
	UUID uuid.UUID `json:"uuid"`
}

type ID struct {
	Id int64
}
type CoffeeSpecDetail struct {
	Name             string                         `json:"name"`
	CoffeeSpecDetail []coffeeModel.CoffeeSpecDetail `json:"coffee_spec_detail"`
}

func AddCoffeeSpec(c *gin.Context) {
	var coffeeSpec coffeeModel.CoffeeSpec
	var detail CoffeeSpecDetail
	_ = c.ShouldBindJSON(&detail)

	err := coffeeSpec.AddCoffeeSpec(detail.Name, detail.CoffeeSpecDetail)
	if err != nil {
		servers.ReportFormat(c, false, "新增失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "新增成功", gin.H{})
	}
}

func UpdateCoffeeSpec(c *gin.Context) {
	var coffeeSpec coffeeModel.CoffeeSpec
	_ = c.ShouldBindJSON(&coffeeSpec)

	err := coffeeSpec.UpdateCoffeeSpec()
	if err != nil {
		servers.ReportFormat(c, false, "更改失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "更改成功", gin.H{})
	}
}

func GetCoffeeSpecList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	var coffeeSpec coffeeModel.CoffeeSpec
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := coffeeSpec.GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, "获取失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"coffeeSpec": list,
			"total":      total,
			"page":       pageInfo.Page,
			"pageSize":   pageInfo.PageSize,
		})
	}
}

func DeleteCoffeeSpec(c *gin.Context) {
	var Id ID
	var coffeeSpec coffeeModel.CoffeeSpec
	_ = c.ShouldBindJSON(&Id)
	err := coffeeSpec.DeleteCoffeeSpec(Id.Id)
	if err != nil {
		servers.ReportFormat(c, false, "删除失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

func GetCoffeeSpecById(c *gin.Context) {
	var Id ID
	var coffeeSpec coffeeModel.CoffeeSpec
	_ = c.ShouldBindJSON(&Id)

	err := coffeeSpec.GetCoffeeById(Id.Id)

	if err != nil {
		servers.ReportFormat(c, false, "获得失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"coffeeSpec": coffeeSpec})
	}
}

func GetCoffeeSpecByCoffeeId(c *gin.Context) {

}
