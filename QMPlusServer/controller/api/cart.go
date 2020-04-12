package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/customerModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CartInfo struct {
	PageInfo modelInterface.PageInfo `json:"pageInfo"`
	UserId   uuid.UUID               `json:"user_id"`
}

func GetCartList(c *gin.Context) {
	var cartInfo CartInfo
	_ = c.ShouldBindJSON(&cartInfo)
	err, list, total := new(customerModel.Cart).GetInfoListByUserId(cartInfo.PageInfo, cartInfo.UserId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"cartList": list,
			"total":    total,
			"page":     cartInfo.PageInfo.Page,
			"pageSize": cartInfo.PageInfo.PageSize,
			"cart":     cartInfo,
		})
	}
}

type CusCart struct {
	UserId   uuid.UUID `json:"user_id"`
	CoffeeId uuid.UUID `json:"coffee_id"`
}

func AddCart(c *gin.Context) {
	var cusCart CusCart
	var cart customerModel.Cart
	_ = c.ShouldBindJSON(&cusCart)
	err := cart.AddCart(cusCart.UserId, cusCart.CoffeeId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败 %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{"cart": cart})
	}

}

func ReduceCart(c *gin.Context) {
	var cusCart CusCart
	var cart customerModel.Cart
	_ = c.ShouldBindJSON(&cusCart)
	err := cart.ReduceCart(cusCart.UserId, cusCart.CoffeeId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("减少失败 %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "减少成功", gin.H{})
	}

}

type UserId struct {
	UserId uuid.UUID
}

func DelCart(c *gin.Context) {
	var userId UserId
	_ = c.ShouldBindJSON(&userId)
	err := new(customerModel.Cart).DeleteCart(userId.UserId)
	if err != nil {
		servers.ReportFormat(c, false, "删除失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

type CartId struct {
	ID      int64 `json:"id"`
	IsCheck bool  `json:"isCheck"`
}

func CheckStatus(c *gin.Context) {
	var cartId CartId
	_ = c.ShouldBindJSON(&cartId)
	err := new(customerModel.Cart).CheckStatus(cartId.ID, cartId.IsCheck)
	if err != nil {
		servers.ReportFormat(c, false, "修改失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}
