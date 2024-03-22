package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
)

type LinksService struct {
}

// CreateLinks 创建Links记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService) CreateLinks(links webcms.Links) (err error) {
	err = global.GVA_DB.Create(&links).Error
	return err
}

// DeleteLinks 删除Links记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService) DeleteLinks(links webcms.Links) (err error) {
	err = global.GVA_DB.Delete(&links).Error
	return err
}

// DeleteLinksByIds 批量删除Links记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService) DeleteLinksByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.Links{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLinks 更新Links记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService) UpdateLinks(links webcms.Links) (err error) {
	err = global.GVA_DB.Save(&links).Error
	return err
}

// GetLinks 根据id获取Links记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService) GetLinks(id uint) (links webcms.Links, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&links).Error
	return
}

// GetLinksInfoList 分页获取Links记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService) GetLinksInfoList(info webcmsReq.LinksSearch) (list []webcms.Links, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Links{})
	var linkss []webcms.Links
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&linkss).Error
	return linkss, total, err
}

// 获取链接
func (linksService *LinksService) GetLinksinfo() (links []webcms.Links, err error) {
	err = global.GVA_DB.Where("enable = ?", 0).Find(&links).Error
	return
}
