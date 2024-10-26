package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
)

type CliSetService struct{}

// CreateCliSet 创建结算设置记录
// Author [yourname](https://github.com/yourname)
func (clisetService *CliSetService) CreateCliSet(cliset *xiao.CliSet) (err error) {
	err = global.GVA_DB.Create(cliset).Error
	return err
}

// DeleteCliSet 删除结算设置记录
// Author [yourname](https://github.com/yourname)
func (clisetService *CliSetService) DeleteCliSet(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliSet{}, "id = ?", ID).Error
	return err
}

// DeleteCliSetByIds 批量删除结算设置记录
// Author [yourname](https://github.com/yourname)
func (clisetService *CliSetService) DeleteCliSetByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliSet{}, "id in ?", IDs).Error
	return err
}

// UpdateCliSet 更新结算设置记录
// Author [yourname](https://github.com/yourname)
func (clisetService *CliSetService) UpdateCliSet(cliset xiao.CliSet) (err error) {
	err = global.GVA_DB.Model(&xiao.CliSet{}).Where("id = ?", cliset.ID).Updates(&cliset).Error
	return err
}

// GetCliSet 根据ID获取结算设置记录
// Author [yourname](https://github.com/yourname)
func (clisetService *CliSetService) GetCliSet(ID string) (cliset xiao.CliSet, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliset).Error
	return
}

// GetCliSetInfoList 分页获取结算设置记录
// Author [yourname](https://github.com/yourname)
func (clisetService *CliSetService) GetCliSetInfoList(info xiaoReq.CliSetSearch) (list []xiao.CliSet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliSet{})
	var clisets []xiao.CliSet
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&clisets).Error
	return clisets, total, err
}
func (clisetService *CliSetService) GetCliSetPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
