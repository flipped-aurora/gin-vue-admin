package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
)

type CliOrderService struct{}

// CreateCliOrder 创建订单详情记录
// Author [yourname](https://github.com/yourname)
func (cliOrderService *CliOrderService) CreateCliOrder(cliOrder *xiao.CliOrder) (err error) {
	err = global.GVA_DB.Create(cliOrder).Error
	return err
}

// DeleteCliOrder 删除订单详情记录
// Author [yourname](https://github.com/yourname)
func (cliOrderService *CliOrderService) DeleteCliOrder(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliOrder{}, "id = ?", ID).Error
	return err
}

// DeleteCliOrderByIds 批量删除订单详情记录
// Author [yourname](https://github.com/yourname)
func (cliOrderService *CliOrderService) DeleteCliOrderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliOrder{}, "id in ?", IDs).Error
	return err
}

// UpdateCliOrder 更新订单详情记录
// Author [yourname](https://github.com/yourname)
func (cliOrderService *CliOrderService) UpdateCliOrder(cliOrder xiao.CliOrder) (err error) {
	err = global.GVA_DB.Model(&xiao.CliOrder{}).Where("id = ?", cliOrder.ID).Updates(&cliOrder).Error
	return err
}

// GetCliOrder 根据ID获取订单详情记录
// Author [yourname](https://github.com/yourname)
func (cliOrderService *CliOrderService) GetCliOrder(ID string) (cliOrder xiao.CliOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliOrder).Error
	return
}

// GetCliOrderInfoList 分页获取订单详情记录
// Author [yourname](https://github.com/yourname)
func (cliOrderService *CliOrderService) GetCliOrderInfoList(info xiaoReq.CliOrderSearch) (list []xiao.CliOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliOrder{})
	var cliOrders []xiao.CliOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Text != "" {
		db = db.Where("text LIKE ?", "%"+info.Text+"%")
	}
	if info.Status != "" {
		db = db.Where("status LIKE ?", "%"+info.Status+"%")
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["amount"] = true
	orderMap["num"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cliOrders).Error
	return cliOrders, total, err
}
func (cliOrderService *CliOrderService) GetCliOrderPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
