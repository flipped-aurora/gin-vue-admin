package webcms

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
)

type SwiperService struct {
}

// CreateSwiper 创建Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) CreateSwiper(swiper webcms.Swiper) (err error) {
	err = global.GVA_DB.Create(&swiper).Error
	return err
}

// DeleteSwiper 删除Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) DeleteSwiper(swiper webcms.Swiper) (err error) {
	err = global.GVA_DB.Delete(&swiper).Error
	return err
}

// DeleteSwiperByIds 批量删除Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) DeleteSwiperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.Swiper{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSwiper 更新Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) UpdateSwiper(swiper webcms.Swiper) (err error) {
	err = global.GVA_DB.Save(&swiper).Error
	return err
}

// GetSwiper 根据id获取Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) GetSwiper(id uint) (swiper webcms.Swiper, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&swiper).Error
	return
}

// GetSwiperInfoList 分页获取Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) GetSwiperInfoList(info webcmsReq.SwiperSearch, siteid string) (list []webcms.Swiper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Swiper{}).Where("siteid", siteid)
	var swipers []webcms.Swiper
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&swipers).Error
	return swipers, total, err
}

// GetSwiperList 获取Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) GetSwiperList(siteid int) (list []webcms.Swiper) {
	// 创建db
	db := global.GVA_DB.Model(&webcms.Swiper{}).Where("siteid", siteid)
	var swipers []webcms.Swiper
	err := db.Where("enable = ?", 1).Find(&swipers).Error
	if err != nil {
		fmt.Println(err)
	}
	return swipers
}
