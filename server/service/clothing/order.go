package clothing

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"gorm.io/gorm"
	"time"
)

type OrderService struct {
}

// CreateOrder 创建Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) CreateOrder(order *clothing.Order) (err error) {
	err = global.GVA_DB.Create(order).Error
	return err
}

// DeleteOrder 删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrder(order clothing.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Order{}).Where("id = ?", order.ID).Update("deleted_by", order.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&order).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteOrderByIds 批量删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrderByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Order{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Order{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateOrder 更新Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) UpdateOrder(order clothing.Order) (err error) {
	err = global.GVA_DB.Save(&order).Error
	return err
}

// GetOrder 根据id获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrder(id uint) (order clothing.Order, err error) {
	err = global.GVA_DB.Preload("User").Preload("Company").Where("id = ?", id).First(&order).Error
	return
}

func (orderService *OrderService) GetOrderByOrderNo(orderNo string) (order clothing.Order, err error) {
	err = global.GVA_DB.Preload("User").Preload("Company").Where("order_no = ?", orderNo).First(&order).Error
	return
}

// GetOrderInfoList 分页获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrderInfoList(info clothingReq.OrderSearch) (list []clothing.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Order{})
	var orders []clothing.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+info.OrderNo+"%")
	}
	if info.PayNo != "" {
		db = db.Where("pay_no LIKE ?", "%"+info.PayNo+"%")
	}
	if info.CompanyID != 0 {
		db = db.Where("company_id = ?", info.CompanyID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("User").Preload("Company").Limit(limit).Offset(offset).Find(&orders).Error
	return orders, total, err
}

func (orderService *OrderService) PaySuccess(order clothing.Order, payNo string, payTime time.Time) (err error) {
	if order.PayStatus != enum.Pending || order.Status != enum.Pending {
		return errors.New("order processed")
	}
	if err := global.GVA_DB.Model(&order).Updates(map[string]interface{}{
		"pay_no":     payNo,
		"pay_at":     payTime,
		"pay_status": enum.Success,
		"status":     enum.Success,
	}).Error; err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return err
	}
	var expirationAt time.Time
	if order.Company.ExpirationAt.Sub(time.Now()).Seconds() < 0 {
		expirationAt = time.Now().AddDate(0, 0, order.Day)
	} else {
		expirationAt = order.Company.ExpirationAt.AddDate(0, 0, order.Day)
	}
	err = global.GVA_DB.Model(&order.Company).Update("expiration_at", expirationAt.Format("2006-01-02")).Error
	return err
}
