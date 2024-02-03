package discuss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/discuss"
    discussReq "github.com/flipped-aurora/gin-vue-admin/server/model/discuss/request"
)

type DisInfoService struct {
}

// CreateDisInfo 创建帖子信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (disDataService *DisInfoService) CreateDisInfo(disData *discuss.DisInfo) (err error) {
	err = global.GVA_DB.Create(disData).Error
	return err
}

// DeleteDisInfo 删除帖子信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (disDataService *DisInfoService)DeleteDisInfo(id string) (err error) {
	err = global.GVA_DB.Delete(&discuss.DisInfo{},"id = ?",id).Error
	return err
}

// DeleteDisInfoByIds 批量删除帖子信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (disDataService *DisInfoService)DeleteDisInfoByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]discuss.DisInfo{},"id in ?",ids).Error
	return err
}

// UpdateDisInfo 更新帖子信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (disDataService *DisInfoService)UpdateDisInfo(disData discuss.DisInfo) (err error) {
	err = global.GVA_DB.Save(&disData).Error
	return err
}

// GetDisInfo 根据id获取帖子信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (disDataService *DisInfoService)GetDisInfo(id string) (disData discuss.DisInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&disData).Error
	return
}

// GetDisInfoInfoList 分页获取帖子信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (disDataService *DisInfoService)GetDisInfoInfoList(info discussReq.DisInfoSearch) (list []discuss.DisInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&discuss.DisInfo{})
    var disDatas []discuss.DisInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.DisTitle != "" {
        db = db.Where("dis_title LIKE ?","%"+ info.DisTitle+"%")
    }
    if info.DisComId != nil {
        db = db.Where("dis_com_id = ?",info.DisComId)
    }
    if info.DisUserId != nil {
        db = db.Where("dis_user_id = ?",info.DisUserId)
    }
    if info.DisModel != nil {
        db = db.Where("dis_model = ?",info.DisModel)
    }
    if info.DisStatus != nil {
        db = db.Where("dis_status = ?",info.DisStatus)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&disDatas).Error
	return  disDatas, total, err
}
