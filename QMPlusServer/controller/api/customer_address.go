package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/customerModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddAddress(c *gin.Context) {
	var address customerModel.Address
	_ = c.BindJSON(&address)
	err := address.AddAddress()
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

func UpdateAddress(c *gin.Context) {
	var address customerModel.Address
	_ = c.BindJSON(&address)
	err := address.UpdateAddress()
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func DeleteAddress(c *gin.Context) {
	var address customerModel.Address
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := address.DeleteAddress(id)
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

func GetAddressList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(customerModel.Address).GetInfoList(pageInfo)
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
