package orderinfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderinfo"
	orderinfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderinfo/request"
)

type OrderInfoService struct {
}

// CreateOrderInfo 创建OrderInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderInfoService *OrderInfoService) CreateOrderInfo(orderInfo *orderinfo.OrderInfo) (err error) {
	err = global.MustGetGlobalDBByDBName("票务服务").Create(orderInfo).Error
	return err
}

// DeleteOrderInfo 删除OrderInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderInfoService *OrderInfoService) DeleteOrderInfo(orderInfo orderinfo.OrderInfo) (err error) {
	err = global.MustGetGlobalDBByDBName("票务服务").Delete(&orderInfo).Error
	return err
}

// DeleteOrderInfoByIds 批量删除OrderInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderInfoService *OrderInfoService) DeleteOrderInfoByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("票务服务").Delete(&[]orderinfo.OrderInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrderInfo 更新OrderInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderInfoService *OrderInfoService) UpdateOrderInfo(orderInfo orderinfo.OrderInfo) (err error) {
	err = global.MustGetGlobalDBByDBName("票务服务").Save(&orderInfo).Error
	return err
}

// GetOrderInfo 根据id获取OrderInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderInfoService *OrderInfoService) GetOrderInfo(id uint) (orderInfo orderinfo.OrderInfo, err error) {
	err = global.MustGetGlobalDBByDBName("票务服务").Where("id = ?", id).First(&orderInfo).Error
	return
}

// GetOrderInfoInfoList 分页获取OrderInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderInfoService *OrderInfoService) GetOrderInfoInfoList(info orderinfoReq.OrderInfoSearch) (list []orderinfo.OrderInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("票务服务").Model(&orderinfo.OrderInfo{})
	var orderInfos []orderinfo.OrderInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrderId != "" {
		db = db.Where("order_id LIKE ?", "%"+info.OrderId+"%")
	}
	if info.TicketNumber != "" {
		db = db.Where("ticket_number LIKE ?", "%"+info.TicketNumber+"%")
	}
	if info.ContactPhone != "" {
		db = db.Where("contact_phone LIKE ?", "%"+info.ContactPhone+"%")
	}
	if info.MachineName != "" {
		db = db.Where("machine_name LIKE ?", "%"+info.MachineName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&orderInfos).Error
	return orderInfos, total, err
}
