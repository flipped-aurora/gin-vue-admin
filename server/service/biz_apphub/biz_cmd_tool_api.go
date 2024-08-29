package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"gorm.io/gorm"
)

type BizCmdToolApiService struct{}

// CreateBizCmdToolApi 创建后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCmdToolApiService *BizCmdToolApiService) CreateBizCmdToolApi(bizCmdToolApi *biz_apphub.BizCmdToolApi) (err error) {
	err = global.GVA_DB.Create(bizCmdToolApi).Error
	return err
}

// DeleteBizCmdToolApi 删除后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCmdToolApiService *BizCmdToolApiService) DeleteBizCmdToolApi(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&biz_apphub.BizCmdToolApi{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&biz_apphub.BizCmdToolApi{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBizCmdToolApiByIds 批量删除后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCmdToolApiService *BizCmdToolApiService) DeleteBizCmdToolApiByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&biz_apphub.BizCmdToolApi{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&biz_apphub.BizCmdToolApi{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBizCmdToolApi 更新后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCmdToolApiService *BizCmdToolApiService) UpdateBizCmdToolApi(bizCmdToolApi biz_apphub.BizCmdToolApi) (err error) {
	err = global.GVA_DB.Model(&biz_apphub.BizCmdToolApi{}).Where("id = ?", bizCmdToolApi.ID).Updates(&bizCmdToolApi).Error
	return err
}

// GetBizCmdToolApi 根据ID获取后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCmdToolApiService *BizCmdToolApiService) GetBizCmdToolApi(ID string) (bizCmdToolApi biz_apphub.BizCmdToolApi, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bizCmdToolApi).Error
	return
}

// GetBizCmdToolApiInfoList 分页获取后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCmdToolApiService *BizCmdToolApiService) GetBizCmdToolApiInfoList(info biz_apphubReq.BizCmdToolApiSearch) (list []biz_apphub.BizCmdToolApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&biz_apphub.BizCmdToolApi{})
	var bizCmdToolApis []biz_apphub.BizCmdToolApi
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
	if info.ApiPath != "" {
		db = db.Where("api_path LIKE ?", "%"+info.ApiPath+"%")
	}
	if info.IsPublic != "" {
		db = db.Where("is_public = ?", info.IsPublic)
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

	err = db.Find(&bizCmdToolApis).Error
	return bizCmdToolApis, total, err
}
