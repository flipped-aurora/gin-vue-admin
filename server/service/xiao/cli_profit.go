package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
)

type CliProfitService struct{}

// CreateCliProfit 创建结算详情记录
// Author [yourname](https://github.com/yourname)
func (cliprofitService *CliProfitService) CreateCliProfit(cliprofit *xiao.CliProfit) (err error) {
	err = global.GVA_DB.Create(cliprofit).Error
	return err
}

// DeleteCliProfit 删除结算详情记录
// Author [yourname](https://github.com/yourname)
func (cliprofitService *CliProfitService) DeleteCliProfit(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliProfit{}, "id = ?", ID).Error
	return err
}

// DeleteCliProfitByIds 批量删除结算详情记录
// Author [yourname](https://github.com/yourname)
func (cliprofitService *CliProfitService) DeleteCliProfitByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliProfit{}, "id in ?", IDs).Error
	return err
}

// UpdateCliProfit 更新结算详情记录
// Author [yourname](https://github.com/yourname)
func (cliprofitService *CliProfitService) UpdateCliProfit(cliprofit xiao.CliProfit) (err error) {
	err = global.GVA_DB.Model(&xiao.CliProfit{}).Where("id = ?", cliprofit.ID).Updates(&cliprofit).Error
	return err
}

// GetCliProfit 根据ID获取结算详情记录
// Author [yourname](https://github.com/yourname)
func (cliprofitService *CliProfitService) GetCliProfit(ID string) (cliprofit xiao.CliProfit, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliprofit).Error
	return
}

// GetCliProfitInfoList 分页获取结算详情记录
// Author [yourname](https://github.com/yourname)
func (cliprofitService *CliProfitService) GetCliProfitInfoList(info xiaoReq.CliProfitSearch) (list []xiao.CliProfit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliProfit{})
	var cliprofits []xiao.CliProfit
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

	err = db.Find(&cliprofits).Error
	return cliprofits, total, err
}
func (cliprofitService *CliProfitService) GetCliProfitPublic() {

	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (cliprofitService *CliProfitService) WithAsset(profit *xiao.CliProfit) (err error) {

	//tx := global.GVA_DB.Begin()
	//1.查询投入了多少资产 2.查询提取了多少资产 3.查询剩余多少资产可提
	return nil

}
