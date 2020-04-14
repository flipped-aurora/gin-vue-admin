package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/customerModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

type Order struct {
	PageInfo  modelInterface.PageInfo `json:",inline"`
	OrderType int                     `json:"orderType"`
	UserId    uuid.UUID               `json:"userId"`
}

func GetOrderListByOrderType(c *gin.Context) {
	var order Order
	_ = c.ShouldBindJSON(&order)
	err, list, total := new(customerModel.CustomerOrder).GetInfoListByOrderType(order.PageInfo, order.OrderType, order.UserId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"orderList": list,
			"total":     total,
			"page":      order.PageInfo.Page,
			"pageSize":  order.PageInfo.PageSize,
		})
	}
}

type CartList struct {
	CartList []customerModel.Cart `json:"cartList"`
}

func AddOrder(c *gin.Context) {
	var cartList CartList
	var order customerModel.CustomerOrder
	_ = c.ShouldBindJSON(&cartList)
	err := order.AddOrder(cartList.CartList)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败, %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

type OrderId struct {
	OrderId uuid.UUID `json:"orderId"`
}

func DeleteOrder(c *gin.Context) {
	var orderId OrderId
	_ = c.ShouldBindJSON(&orderId)
	err := new(customerModel.CustomerOrder).DeleteOrder(orderId.OrderId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败, %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}
