package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
)

type CliMainprofitService struct{}

// CreateCliMainprofit 创建结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) CreateCliMainprofit(climainprofit *xiao.CliMainprofit) (err error) {
	err = global.GVA_DB.Create(climainprofit).Error
	return err
}

// DeleteCliMainprofit 删除结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) DeleteCliMainprofit(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliMainprofit{}, "id = ?", ID).Error
	return err
}

// DeleteCliMainprofitByIds 批量删除结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) DeleteCliMainprofitByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliMainprofit{}, "id in ?", IDs).Error
	return err
}

// UpdateCliMainprofit 更新结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) UpdateCliMainprofit(climainprofit xiao.CliMainprofit) (err error) {
	err = global.GVA_DB.Model(&xiao.CliMainprofit{}).Where("id = ?", climainprofit.ID).Updates(&climainprofit).Error
	return err
}

// GetCliMainprofit 根据ID获取结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) GetCliMainprofit(ID string) (climainprofit xiao.CliMainprofit, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&climainprofit).Error
	return
}

// GetCliMainprofitInfoList 分页获取结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) GetCliMainprofitInfoList(info xiaoReq.CliMainprofitSearch) (list []xiao.CliMainprofit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliMainprofit{})
	var climainprofits []xiao.CliMainprofit
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
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
	orderMap["static"] = true
	orderMap["indirect"] = true
	orderMap["team"] = true
	orderMap["amount"] = true
	orderMap["descnum"] = true
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

	err = db.Find(&climainprofits).Error
	return climainprofits, total, err
}
func (climainprofitService *CliMainprofitService) GetCliMainprofitPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
