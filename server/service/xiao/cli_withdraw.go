package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
)

type CliWithdrawService struct{}

// CreateCliWithdraw 创建提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) CreateCliWithdraw(cliwithdraw *xiao.CliWithdraw) (err error) {
	err = global.GVA_DB.Create(cliwithdraw).Error
	return err
}

// DeleteCliWithdraw 删除提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) DeleteCliWithdraw(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliWithdraw{}, "id = ?", ID).Error
	return err
}

// DeleteCliWithdrawByIds 批量删除提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) DeleteCliWithdrawByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliWithdraw{}, "id in ?", IDs).Error
	return err
}

// UpdateCliWithdraw 更新提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) UpdateCliWithdraw(cliwithdraw xiao.CliWithdraw) (err error) {
	err = global.GVA_DB.Model(&xiao.CliWithdraw{}).Where("id = ?", cliwithdraw.ID).Updates(&cliwithdraw).Error
	return err
}

// GetCliWithdraw 根据ID获取提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) GetCliWithdraw(ID string) (cliwithdraw xiao.CliWithdraw, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliwithdraw).Error
	return
}

// GetCliWithdrawInfoList 分页获取提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) GetCliWithdrawInfoList(info xiaoReq.CliWithdrawSearch) (list []xiao.CliWithdraw, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliWithdraw{})
	var cliwithdraws []xiao.CliWithdraw
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

	err = db.Find(&cliwithdraws).Error
	return cliwithdraws, total, err
}
func (cliwithdrawService *CliWithdrawService) GetCliWithdrawPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
