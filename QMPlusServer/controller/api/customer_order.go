package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/customerModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
)

func GetOrderList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(customerModel.CustomerOrder).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"orderList": list,
			"total":     total,
			"page":      pageInfo.Page,
			"pageSize":  pageInfo.PageSize,
		})
	}
}

func AddOrder(c *gin.Context) {
	var order customerModel.CustomerOrder
	_ = c.ShouldBindJSON(&order)
	err := order.AddOrder()
	if err != nil {
		servers.ReportFormat(c, false, "添加失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

func UpdateOrder(c *gin.Context) {
	var order customerModel.CustomerOrder
	err := order.UpdateOrder()
	if err != nil {
		servers.ReportFormat(c, false, "修改失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

type OrderId struct {
	Id int64
}

func DeleteOrder(c *gin.Context) {
	var order customerModel.CustomerOrder
	var orderId OrderId
	_ = c.ShouldBindJSON(&orderId)
	err := order.DeleteOrder(orderId.Id)
	if err != nil {
		servers.ReportFormat(c, false, "删除失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}
