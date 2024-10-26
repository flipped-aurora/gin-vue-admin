package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
)

type CliSetvipService struct{}

// CreateCliSetvip 创建团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) CreateCliSetvip(clisetvip *xiao.CliSetvip) (err error) {
	err = global.GVA_DB.Create(clisetvip).Error
	return err
}

// DeleteCliSetvip 删除团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) DeleteCliSetvip(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliSetvip{}, "id = ?", ID).Error
	return err
}

// DeleteCliSetvipByIds 批量删除团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) DeleteCliSetvipByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliSetvip{}, "id in ?", IDs).Error
	return err
}

// UpdateCliSetvip 更新团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) UpdateCliSetvip(clisetvip xiao.CliSetvip) (err error) {
	err = global.GVA_DB.Model(&xiao.CliSetvip{}).Where("id = ?", clisetvip.ID).Updates(&clisetvip).Error
	return err
}

// GetCliSetvip 根据ID获取团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) GetCliSetvip(ID string) (clisetvip xiao.CliSetvip, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&clisetvip).Error
	return
}

// GetCliSetvipInfoList 分页获取团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) GetCliSetvipInfoList(info xiaoReq.CliSetvipSearch) (list []xiao.CliSetvip, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliSetvip{})
	var clisetvips []xiao.CliSetvip
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

	err = db.Find(&clisetvips).Error
	return clisetvips, total, err
}
func (clisetvipService *CliSetvipService) GetCliSetvipPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
