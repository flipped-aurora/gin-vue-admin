package shop

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	alipay "github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/service"
	wxpay "github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type ShopOrdersService struct {
}

// CreateShopOrders 创建shopOrders表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopOrdersService *ShopOrdersService) CreateShopOrders(shopOrders *shop.ShopOrders) (err error) {
	err = global.GVA_DB.Create(shopOrders).Error
	return err
}

// DeleteShopOrders 删除shopOrders表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopOrdersService *ShopOrdersService) DeleteShopOrders(shopOrders shop.ShopOrders) (err error) {
	err = global.GVA_DB.Delete(&shopOrders).Error
	return err
}

// DeleteShopOrdersByIds 批量删除shopOrders表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopOrdersService *ShopOrdersService) DeleteShopOrdersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]shop.ShopOrders{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateShopOrders 更新shopOrders表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopOrdersService *ShopOrdersService) UpdateShopOrders(shopOrders shop.ShopOrders) (err error) {
	err = global.GVA_DB.Save(&shopOrders).Error
	return err
}

// GetShopOrders 根据id获取shopOrders表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopOrdersService *ShopOrdersService) GetShopOrders(id uint) (shopOrders shop.ShopOrders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shopOrders).Error
	return
}

// RefundShopOrders 根据OutTradeNo退款
// Author [piexlmax](https://github.com/piexlmax)
func (shopOrdersService *ShopOrdersService) RefundShopOrders(tradeNo string) (order *shop.ShopOrders, err error) {
	//从数据库把商品信息取出来
	var shopOrders *shop.ShopOrders
	err = global.GVA_DB.Where("OutTradeNo = ? ", tradeNo).First(&shopOrders).Error
	if err != nil {
		return nil, errors.New("商品信息错误")
	}
	if shopOrders.PayMent == utils.Wxpay { //微信退款
		order, err = wxpay.ServiceGroupApp.PlugServerRefunds(tradeNo)
	} else if shopOrders.PayMent == utils.Alipay { //支付宝退款
		order, err = alipay.ServiceGroupApp.PlugServiceRefund(tradeNo)
	} else {
		return nil, errors.New("请求参数有误：" + tradeNo)
	}
	if err != nil {
		return nil, err
	}
	if order.TradeState != utils.REFUND {
		return nil, errors.New("更新数据库失败,订单号：" + tradeNo)
	}
	return order, nil
}

// GetShopOrdersInfoList 分页获取shopOrders表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopOrdersService *ShopOrdersService) GetShopOrdersInfoList(info shopReq.ShopOrdersSearch) (list []shop.ShopOrders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&shop.ShopOrders{})
	var shopOrderss []shop.ShopOrders
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AppId != "" {
		db = db.Where("AppId = ?", info.AppId)
	}
	if info.MchId != "" {
		db = db.Where("MchId = ?", info.MchId)
	}
	if info.OutTradeNo != "" {
		db = db.Where("OutTradeNo = ?", info.OutTradeNo)
	}
	if info.TransactionId != "" {
		db = db.Where("TransactionId = ?", info.TransactionId)
	}
	if info.TradeState != "" {
		db = db.Where("TradeState = ?", info.TradeState)
	}
	//if info.OpenId != "" {
	//	db = db.Where("OpenId = ?", info.OpenId)
	//}
	//if info.Total != nil {
	//	db = db.Where("Total = ?", info.Total)
	//}
	if info.PayerTotal != nil {
		db = db.Where("PayerTotal = ?", info.PayerTotal)
	}
	if info.GoodsTitle != "" {
		db = db.Where("GoodsTitle LIKE ?", "%"+info.GoodsTitle+"%")
	}
	if info.PayMent != "" {
		db = db.Where("PayMent = ?", info.PayMent)
	}
	if info.DeviceId != "" {
		db = db.Where("DeviceId = ?", info.DeviceId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Order("id desc").Limit(limit).Offset(offset)
	}
	err = db.Find(&shopOrderss).Error
	return shopOrderss, total, err
}

// getUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
func getUserInfo(c *gin.Context) uint {
	uuid := utils.GetUserUuid(c)
	var user system.UserService
	ReqUser, err := user.GetUserInfo(uuid)
	if err != nil {
		return 0
	}
	return ReqUser.ID
}
