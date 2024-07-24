package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
)

var Info = new(info)

type info struct{}

// CreateInfo 创建公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *info) CreateInfo(info *model.Info) (err error) {
	err = global.GVA_DB.Create(info).Error
	return err
}

// DeleteInfo 删除公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *info) DeleteInfo(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Info{}, "id = ?", ID).Error
	return err
}

// DeleteInfoByIds 批量删除公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *info) DeleteInfoByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Info{}, "id in ?", IDs).Error
	return err
}

// UpdateInfo 更新公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *info) UpdateInfo(info model.Info) (err error) {
	err = global.GVA_DB.Model(&model.Info{}).Where("id = ?", info.ID).Updates(&info).Error
	return err
}

// GetInfo 根据ID获取公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *info) GetInfo(ID string) (info model.Info, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&info).Error
	return
}

// GetInfoInfoList 分页获取公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *info) GetInfoInfoList(info request.InfoSearch) (list []model.Info, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Info{})
	var infos []model.Info
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
	err = db.Find(&infos).Error
	return infos, total, err
}
func (s *info) GetInfoDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userID := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_users").Select("nick_name as label,id as value").Scan(&userID)
	res["userID"] = userID
	return
}
