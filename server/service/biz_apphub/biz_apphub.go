package biz_apphub

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"gorm.io/gorm"
	"time"
)

type BizAppHubService struct{}

// CreateBizAppHub 创建biz_apphub记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizAppHubService *BizAppHubService) CreateBizAppHub(bizAppHub *biz_apphub.BizAppHub) (err error) {
	index, err := bizAppHubService.Deploy(*bizAppHub)
	if err != nil {
		return err
	}
	bizAppHub.IndexHtml = index
	err = global.GVA_DB.Create(bizAppHub).Error
	if err != nil {
		return err
	}

	record := *bizAppHub
	record.CreatedAt = time.Now()
	record.ID = 0

	record.OperateUser = bizAppHub.OperateUser
	bz := &biz_apphub.BizAppHubRecord{
		AppId:     bizAppHub.ID,
		BizAppHub: record,
		//UpdatedUser: bizAppHub.OperateUser,
	}
	err = global.GVA_DB.Create(bz).Error
	if err != nil {
		return err
	}
	return err
}

// DeleteBizAppHub 删除biz_apphub记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizAppHubService *BizAppHubService) DeleteBizAppHub(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&biz_apphub.BizAppHub{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&biz_apphub.BizAppHub{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBizAppHubByIds 批量删除biz_apphub记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizAppHubService *BizAppHubService) DeleteBizAppHubByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&biz_apphub.BizAppHub{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&biz_apphub.BizAppHub{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBizAppHub 更新biz_apphub记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizAppHubService *BizAppHubService) UpdateBizAppHub(bizAppHub biz_apphub.BizAppHub) (err error) {

	var b biz_apphub.BizAppHub
	err = global.GVA_DB.Model(&biz_apphub.BizAppHub{}).Where("id = ?", bizAppHub.ID).First(&b).Error
	if err != nil {
		return err
	}
	if bizAppHub.Version != b.Version {
		//更新版本
		if bizAppHub.OssPath == "" {
			return fmt.Errorf("文件地址不能为空")
		}
		index, err := bizAppHubService.Deploy(bizAppHub)
		if err != nil {
			return err
		}
		bizAppHub.IndexHtml = index
		//b.ID = 0
		//b.CreatedAt = time.Now()

		bizAppHubRecord := bizAppHub
		bizAppHubRecord.ID = 0
		bizAppHubRecord.CreatedAt = time.Now()
		bizAppHubRecord.IndexHtml = index
		//记录更新版本
		err = global.GVA_DB.Model(&biz_apphub.BizAppHubRecord{}).
			Create(&biz_apphub.BizAppHubRecord{AppId: bizAppHub.ID,
				BizAppHub: bizAppHubRecord}).Error
		if err != nil {
			return err
		}
	}

	err = global.GVA_DB.Model(&biz_apphub.BizAppHub{}).
		Where("id = ?", bizAppHub.ID).Updates(&bizAppHub).Error
	return err
}

// GetBizAppHub 根据ID获取biz_apphub记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizAppHubService *BizAppHubService) GetBizAppHub(ID string) (bizAppHub biz_apphub.BizAppHub, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bizAppHub).Error
	return
}

// GetBizAppHub 根据ID获取biz_apphub记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizAppHubService *BizAppHubService) GetBizAppHubRecord(ID string) (bizAppHub biz_apphub.BizAppHub, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bizAppHub).Error
	return
}

// GetBizAppHubInfoList 分页获取biz_apphub记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizAppHubService *BizAppHubService) GetBizAppHubInfoList(info biz_apphubReq.BizAppHubSearch) (list []biz_apphub.BizAppHub, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&biz_apphub.BizAppHub{})
	var bizAppHubs []biz_apphub.BizAppHub
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AppName != "" {
		db = db.Where("app_name = ?", info.AppName)
	}
	if info.AppCode != "" {
		db = db.Where("app_code = ?", info.AppCode)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	if info.Classify != "" {
		db = db.Where("classify = ?", info.Classify)
	}
	if info.Version != "" {
		db = db.Where("version = ?", info.Version)
	}
	if info.Mode != "" {
		db = db.Where("mode = ?", info.Mode)
	}
	if info.Tags != "" {
		db = db.Where("tags LIKE ?", "%"+info.Tags+"%")
	}
	if info.Video != "" {
		db = db.Where("video = ?", info.Video)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["version"] = true
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

	err = db.Find(&bizAppHubs).Error
	return bizAppHubs, total, err
}
